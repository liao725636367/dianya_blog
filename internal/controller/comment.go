package controller

import (
	"blog/apiv1"
	"blog/internal/service"
	"context"
)

type cComment struct {
}

var Comment = cComment{}

// Add 增加评论
func (c *cComment) Add(ctx context.Context, in *apiv1.CommentAddReq) (out *apiv1.CommentAddRes, err error) {
	return nil, service.Comment().Add(ctx, in)
}

// List 评论列表
func (c *cComment) List(ctx context.Context, in *apiv1.CommentListReq) (out *apiv1.CommentListRes, err error) {
	count, list, err := service.Comment().List(ctx, in)
	if err != nil {
		return nil, err
	}
	out = &apiv1.CommentListRes{
		Count: count,
		List:  make([]*apiv1.CommentRow, 0),
	}
	for _, v := range list {
		out.List = append(out.List, &apiv1.CommentRow{
			Id:        v.Id,
			Content:   v.Content,
			Nickname:  v.Nickname,
			ArticleId: v.ArticleId,
			CreatedAt: v.CreatedAt,
		})
	}
	return out, nil
}
