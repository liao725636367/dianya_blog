package service

import (
	"blog/apiv1"
	"blog/internal/model/entity"
	"blog/internal/service/internal/dao"
	"blog/utility/utils"
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type IComment interface {
	// List 评论列表
	List(ctx context.Context, in *apiv1.CommentListReq) (int, []*entity.BlogComment, error)
	// Add 添加评论
	Add(ctx context.Context, in *apiv1.CommentAddReq) error
}
type (
	sComment struct{}
)

var insComment = sComment{}

func Comment() IComment {
	return &insComment
}

//获取评论列表
func (c *sComment) List(ctx context.Context, in *apiv1.CommentListReq) (int, []*entity.BlogComment, error) {
	//结构体校验
	err1 := g.Validator().Data(in).Run(ctx)
	if err1 != nil {
		return 0, nil, err1.FirstError()
	}
	if in.ArticleId == 0 {
		return 0, nil, utils.CodeErr(1, "只能获取文章评论")
	}
	list := make([]*entity.BlogComment, 0)
	count, err := dao.BlogComment.Ctx(ctx).Count()
	if err != nil {
		return 0, nil, err
	}
	err = dao.BlogComment.Ctx(ctx).Where("article_id=?", in.ArticleId).Page(in.Page, in.Size).Scan(&list)
	if err != nil {
		return 0, nil, err
	}
	return count, list, nil
}

//增加评论
func (c *sComment) Add(ctx context.Context, in *apiv1.CommentAddReq) error {
	//结构体校验
	err1 := g.Validator().Data(in).Run(ctx)
	if err1 != nil {
		return err1.FirstError()
	}
	//检查文章是否存在
	_, err := Article().Row(ctx, in.ArticleId)
	if err != nil {
		return err
	}
	_, err = dao.BlogComment.Ctx(ctx).Insert(&entity.BlogComment{
		Content:   in.Content,
		Nickname:  in.Nickname,
		Email:     in.Email,
		ArticleId: in.ArticleId,
	})
	return err
}
