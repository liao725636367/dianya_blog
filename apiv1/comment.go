package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type CommentListReq struct {
	g.Meta    `path:"/article/{id}/comments" method:"get" tags:"评论" summary:"获取文章评论列表"`
	ArticleId uint `json:"id"`
	Page      int  `json:"page"`
	Size      int  `json:"size"`
}
type CommentRow struct {
	Id        uint        `json:"id"        ` // 评论id
	Content   string      `json:"content"   ` // 评论内容
	Nickname  string      `json:"nickname"  ` // 评论者名称
	ArticleId uint        `json:"articleId" ` // 文章id
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
}
type CommentListRes struct {
	Count int           `json:"count"`
	List  []*CommentRow `json:"list"`
}

type CommentAddReq struct {
	g.Meta    `path:"/comment" method:"post" tags:"评论" summary:"添加评论"`
	Content   string `json:"content" v:"required|length:1,100"   ` // 评论内容
	Nickname  string `json:"nickname" v:"required|length:1,100" `  // 评论者名称
	Email     string `json:"email"  v:"required|email"   `         // 评论者邮箱
	ArticleId uint   `json:"articleId" v:"required" `              // 文章id
}
type CommentAddRes struct{}
