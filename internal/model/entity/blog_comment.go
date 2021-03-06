// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-07-08 20:05:29
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BlogComment is the golang structure for table blog_comment.
type BlogComment struct {
	Id        uint        `json:"id"        ` // 评论id
	Content   string      `json:"content"   ` // 评论内容
	Nickname  string      `json:"nickname"  ` // 评论者名称
	Email     string      `json:"email"     ` // 评论者邮箱
	ArticleId uint        `json:"articleId" ` // 文章id
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" ` // 删除时间
}
