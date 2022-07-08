package controller

import (
	"blog/apiv1"
	"blog/internal/service"
	"context"
)

type cCategory struct {
}

var Category = cCategory{}

// List 分类列表
func (c *cCategory) List(ctx context.Context, in *apiv1.CategoryListReq) (out *apiv1.CategoryListRes, err error) {
	list, err := service.Category().GetAllCategory(ctx)
	if err != nil {
		return nil, err
	}
	out = &apiv1.CategoryListRes{List: list}
	return out, nil
}
