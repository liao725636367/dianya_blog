package service

import (
	"blog/apiv1"
	"blog/internal/consts"
	"blog/internal/model"
	"blog/internal/model/entity"
	"blog/internal/service/internal/dao"
	"blog/utility/utils"
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/database/gdb"

	"github.com/gogf/gf/v2/frame/g"
)

type IArticle interface {
	// Add 添加文章
	Add(ctx context.Context, in *apiv1.ArticleAddReq) error
	// Edit 编辑文章
	Edit(ctx context.Context, in *apiv1.ArticleEditReq) error
	// List 文章列表
	List(ctx context.Context, opt *apiv1.ArticleListReq) (int, []*model.Article, error)
	// Row 文章详情
	Row(ctx context.Context, articleId uint, firstVisit ...bool) (*model.Article, error)
	// CheckTitleRepeat 检查文章重复标题
	CheckTitleRepeat(ctx context.Context, title string, articleId ...uint) error
	// AddReduceSupport 文章增加减少点赞数
	AddReduceSupport(ctx context.Context, articleId uint, num int) (newVal int, err error)
}

type (
	sArticle struct{}
)

var insArticle = sArticle{}

func Article() IArticle {
	return &insArticle
}

type ArticleListOpt struct {
	Page    int
	Size    int
	Keyword string
}

// CheckTitleRepeat 标题重复检查
func (s *sArticle) CheckTitleRepeat(ctx context.Context, title string, articleId ...uint) error {
	//检查有没有标题重复文章
	model := dao.BlogArticle.Ctx(ctx).Where("title=?", title)
	if len(articleId) > 0 {
		model = model.Where("id!=?", articleId[0])
	}
	count, err := model.Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return utils.CodeErr(1, "文章标题重复")
	}
	return nil
}

// Add 发布文章
func (s *sArticle) Add(ctx context.Context, in *apiv1.ArticleAddReq) error {
	//结构体校验
	err1 := g.Validator().Data(in).Run(ctx)
	if err1 != nil {
		return err1.FirstError()
	}
	_, err := dao.BlogArticle.Ctx(ctx).Insert(&entity.BlogArticle{
		Author:     fmt.Sprintf("%s", ctx.Value("username")),
		Cover:      in.Cover,
		Title:      in.Title,
		Summary:    in.Summary,
		Content:    in.Content,
		Visit:      0,
		Support:    0,
		CategoryId: 0,
	})
	return err
}

// Edit 修改文章
func (s *sArticle) Edit(ctx context.Context, in *apiv1.ArticleEditReq) error {
	//结构体校验
	err1 := g.Validator().Data(in).Run(ctx)
	if err1 != nil {
		return err1.FirstError()
	}
	articleId := in.Id
	if err := s.CheckTitleRepeat(ctx, in.Title, articleId); err != nil {
		return err
	}
	//检查有没有此文章
	article, err := s.Row(ctx, articleId)
	if err != nil {
		return err
	}

	if article == nil {
		return utils.CodeErr(1, "待编辑文章不存在")
	}
	article.Cover = in.Cover
	article.Title = in.Title
	article.Summary = in.Summary
	article.Content = in.Content
	article.CategoryId = in.CategoryId
	_, err = dao.BlogArticle.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     fmt.Sprintf(consts.RedisArticleRow, articleId),
		Force:    false,
	}).Where("id=?", articleId).Update(article)
	return err
}

func (s *sArticle) List(ctx context.Context, in *apiv1.ArticleListReq) (int, []*model.Article, error) {
	//结构体校验
	err1 := g.Validator().Data(in).Run(ctx)
	if err1 != nil {
		return 0, nil, err1.FirstError()
	}
	page := in.Page
	size := in.Size
	skip := (page - 1) * size
	table := dao.BlogArticle.Table()
	//if in.Keyword == "" {
	//	return 0, nil, utils.CodeErr(1, "关键词不能为空")
	//}

	fields := "id,author,cover,title,summary,visit,support,category_id,created_at,updated_at"
	option := g.Slice{}
	whereSql := ""
	if in.Keyword != "" {
		whereSql = "where match(title,content) against( ? in  NATURAL LANGUAGE MODE)  "
		option = append(option, in.Keyword)
	}
	sql := fmt.Sprintf("SELECT %s from %s    %s LIMIT %d offset %d", fields, table, whereSql, size, skip)
	countSql := fmt.Sprintf("SELECT * from %s  %s ", table, whereSql)
	count, err := g.DB().GetCount(ctx, countSql, option)
	if err != nil {
		return 0, nil, err
	}
	listResp, err := g.DB().GetAll(ctx, sql, option)
	if err != nil {
		return 0, nil, err
	}
	articleList := make([]*model.Article, 0)
	err = listResp.Structs(&articleList)
	if err != nil {
		return 0, nil, err
	}
	list, err := s.HandleData(ctx, articleList)
	if err != nil {
		return count, nil, err
	}
	return count, list, nil
}

//获取单个文章
func (s *sArticle) Row(ctx context.Context, articleId uint, firstVisit ...bool) (*model.Article, error) {
	var article *model.Article
	err := dao.BlogArticle.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Hour,
		Name:     fmt.Sprintf(consts.RedisArticleRow, articleId),
		Force:    false,
	}).Where("id=?", articleId).Scan(&article)
	if article == nil {
		return nil, utils.CodeErr(1, "文章不存在")
	}
	//增加访问次数
	if len(firstVisit) > 0 && firstVisit[0] {
		_, err = Cache().VisitIncr(ctx, articleId, 1)

	}
	if err != nil {
		return nil, err
	}
	list, err := s.HandleData(ctx, []*model.Article{article})
	if err != nil {
		return nil, err
	}

	return list[0], err
}

//文章数据内部处理
func (s *sArticle) HandleData(ctx context.Context, list []*model.Article) ([]*model.Article, error) {
	articleIds := make([]uint, 0)
	if len(list) == 0 {
		return list, nil
	}
	//获取所有分类
	categoryList, err := Category().GetAllCategory(ctx)
	if err != nil {
		return nil, err
	}

	for _, i2 := range list {
		articleIds = append(articleIds, i2.Id)
		for _, category := range categoryList {
			if category.Id == i2.CategoryId {
				i2.CategoryName = category.Name
			}
		}
	}
	//点赞未刷新数
	supportInfo, err := Cache().GetArticleRedisSupport(ctx, articleIds)
	if err != nil {
		return nil, err
	}
	//阅读未刷新数
	visitInfo, err := Cache().GetArticleRedisVisit(ctx, articleIds)
	if err != nil {
		return nil, err
	}
	for _, article := range list {
		for _, item := range supportInfo {
			if item.Num > 0 && article.Id == item.Id {
				article.Support += uint(item.Num)
			}
		}
		for _, item1 := range visitInfo {
			if item1.Num > 0 && article.Id == item1.Id {
				article.Visit += uint(item1.Num)
			}
		}
	}

	return list, nil
}

//减少或者增加点赞
func (s *sArticle) AddReduceSupport(ctx context.Context, articleId uint, num int) (newVal int, err error) {
	return Cache().SupportIncr(ctx, articleId, num)
}
