package apiv1

import (
	"blog/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type CategoryListReq struct {
	g.Meta `path:"/categories" method:"get" tags:"分类" summary:"获取分类列表"`
}
type CategoryListRes struct {
	List []*entity.BlogCategory `json:"list"`
}
