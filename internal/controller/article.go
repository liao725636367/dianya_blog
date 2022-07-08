package controller

import (
	"blog/apiv1"
	"blog/internal/service"
	"blog/utility/utils"
	"context"

	"github.com/gogf/gf/v2/util/gvalid"
)

type cArticle struct {
}

var Article = cArticle{}

// List 获取文章列表
func (c *cArticle) List(ctx context.Context, in *apiv1.ArticleListReq) (out *apiv1.ArticleListRes, err error) {
	count, list, err := service.Article().List(ctx, in)
	if err != nil {
		return nil, err
	}
	out = &apiv1.ArticleListRes{
		Count: count,
		List:  make([]*apiv1.ArticleRow, 0),
	}
	for _, v := range list {
		out.List = append(out.List, &apiv1.ArticleRow{
			Id:           v.Id,
			Author:       v.Author,
			Cover:        v.Cover,
			Title:        v.Title,
			Summary:      v.Summary,
			Content:      v.Content,
			Visit:        v.Visit,
			Support:      v.Support,
			CategoryName: v.CategoryName,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
		})
	}
	return
}
func (c *cArticle) Row(ctx context.Context, in *apiv1.ArticleReq) (out *apiv1.ArticleRes, err error) {

	v, err := service.Article().Row(ctx, in.Id, in.FirstVisit)
	if err != nil {
		return nil, err
	}
	//fmt.Println("文章id", in.Id)
	out = &apiv1.ArticleRes{
		Id:           v.Id,
		Author:       v.Author,
		Cover:        v.Cover,
		Title:        v.Title,
		Summary:      v.Summary,
		Content:      v.Content,
		Visit:        v.Visit,
		Support:      v.Support,
		CategoryName: v.CategoryName,
		CreatedAt:    v.CreatedAt,
		UpdatedAt:    v.UpdatedAt,
	}
	return out, nil
}

// Add 添加文章
func (c *cArticle) Add(ctx context.Context, in *apiv1.ArticleAddReq) (out *apiv1.ArticleAddRes, err error) {
	if !service.User().IsCtxAdmin(ctx) {
		return nil, utils.CodeErr(403, "无增加文章权限")
	}
	return nil, service.Article().Add(ctx, in)
}

// Edit 编辑文章
func (c *cArticle) Edit(ctx context.Context, in *apiv1.ArticleEditReq) (out *apiv1.ArticleEditRes, err error) {
	if !service.User().IsCtxAdmin(ctx) {
		return nil, utils.CodeErr(403, "无编辑文章权限")
	}
	return nil, service.Article().Edit(ctx, in)
}

// Support 文章点赞
func (c *cArticle) Support(ctx context.Context, in *apiv1.SupportArticleReq) (out *apiv1.SupportArticleRes, err error) {
	_, err = service.Article().AddReduceSupport(ctx, in.Id, 1)
	if err != nil {
		return nil, err
	}
	out = &apiv1.SupportArticleRes{}
	return out, err
}

// UnSupport 文章取消点赞
func (c *cArticle) UnSupport(ctx context.Context, in *apiv1.UnSupportArticleReq) (out *apiv1.UnSupportArticleRes, err error) {
	_, err = service.Article().AddReduceSupport(ctx, in.Id, -1)
	if err != nil {
		return nil, err
	}
	out = &apiv1.UnSupportArticleRes{}
	return out, err
}

func init() {
	//校验规则增加
	gvalid.RegisterRule("titleRepeat", func(ctx context.Context, in gvalid.RuleFuncInput) error {
		title := in.Value.String()
		return service.Article().CheckTitleRepeat(ctx, title)
	})
	gvalid.RegisterRule("hasCategory", func(ctx context.Context, in gvalid.RuleFuncInput) error {
		categoryId := in.Value.Uint()
		categoryList, err := service.Category().GetAllCategory(ctx)
		if err != nil {
			return err
		}
		has := false
	End:
		for _, category := range categoryList {
			if category.Id == categoryId {
				has = true
				break End
			}
		}
		if has {
			return nil
		} else {
			return utils.CodeErr(1, "分类不存在")
		}
	})
}
