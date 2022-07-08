package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type ArticleListReq struct {
	g.Meta  `path:"/articles" method:"get" tags:"文章" summary:"获取文章列表"`
	Page    int    `json:"page" d:"1" v:"min:1"`
	Size    int    `json:"size" d:"10" v:"in:5,10,20"`
	Keyword string `json:"keyword"`
}
type ArticleRow struct {
	Id           uint        `json:"id"         `   // 文章id
	Author       string      `json:"author"     `   // 作者
	Cover        string      `json:"cover"      `   // 文章封面
	Title        string      `json:"title"      `   // 文章标题
	Summary      string      `json:"summary"    `   // 文章摘要
	Content      string      `json:"content"    `   // 文章内容
	Visit        uint        `json:"visit"       `  // 文章阅读量
	Support      uint        `json:"support"    `   // 文章点赞量
	CategoryName string      `json:"categoryName" ` // 文章分类名称
	CreatedAt    *gtime.Time `json:"createdAt"  `   // 创建时间
	UpdatedAt    *gtime.Time `json:"updatedAt"  `   // 更新时间
}
type ArticleListRes struct {
	Count int           `json:"count"`
	List  []*ArticleRow `json:"list"`
}
type ArticleReq struct {
	g.Meta     `path:"/article/{id}" method:"get" tags:"文章" summary:"获取文章详情"`
	Id         uint
	FirstVisit bool `json:"firstVisit"` //是否首次访问页面
}
type ArticleRes ArticleRow
type ArticleAddReq struct {
	g.Meta     `path:"/article" method:"post" tags:"文章" summary:"添加文章"`
	Cover      string `json:"cover" v:"required|length:1,100"`             // 文章封面
	Title      string `json:"title" v:"required|titleRepeat|length:1,100"` // 文章标题
	Summary    string `json:"summary" v:"required|length:1,100"`           // 文章摘要
	CategoryId uint   `json:"categoryId" v:"required|hasCategory"`         // 文章分类id
	Content    string `json:"content" v:"required|length:1,10000"`         // 文章内容
}
type ArticleAddRes struct{}

type ArticleEditReq struct {
	g.Meta     `path:"/article/{id}" method:"put" tags:"文章" summary:"修改文章"`
	Id         uint   `json:"id" v:"required"`
	Cover      string `json:"cover" v:"required|length:1,100"`     // 文章封面
	Title      string `json:"title" v:"required|length:1,100"`     // 文章标题
	Summary    string `json:"summary" v:"required|length:1,100"`   // 文章摘要
	CategoryId uint   `json:"categoryId" v:"required|hasCategory"` // 文章分类id
	Content    string `json:"content" v:"required|length:1,10000"` // 文章内容
}
type ArticleEditRes struct{}

type SupportArticleReq struct {
	g.Meta `path:"/article/support/{id}" method:"post" tags:"文章" summary:"点赞"`
	Id     uint
}
type SupportArticleRes struct{}
type UnSupportArticleReq struct {
	g.Meta `path:"/article/un-support/{id}" method:"post" tags:"文章" summary:"取消点赞"`
	Id     uint
}
type UnSupportArticleRes struct{}
