package service

import (
	"blog/internal/consts"
	"blog/internal/model/entity"
	"blog/internal/service/internal/dao"
	"context"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
)

type ICategory interface {
	GetAllCategory(ctx context.Context) ([]*entity.BlogCategory, error)
}

type (
	sCategory struct{}
)

var insCategory = sCategory{}

func Category() ICategory {
	return &sCategory{}
}

// GetAllCategory 获取全部文章分类
func (c *sCategory) GetAllCategory(ctx context.Context) ([]*entity.BlogCategory, error) {
	var categories = make([]*entity.BlogCategory, 0)
	err := dao.BlogCategory.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Hour * 24,
		Name:     consts.RedisCategory,
		Force:    false,
	}).Scan(&categories)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
