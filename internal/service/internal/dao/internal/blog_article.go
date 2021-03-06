// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-07-08 20:05:29
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BlogArticleDao is the data access object for table blog_article.
type BlogArticleDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns BlogArticleColumns // columns contains all the column names of Table for convenient usage.
}

// BlogArticleColumns defines and stores column names for table blog_article.
type BlogArticleColumns struct {
	Id         string // 文章id
	Author     string // 作者
	Cover      string // 文章封面
	Title      string // 文章标题
	Summary    string // 文章摘要
	Content    string // 文章内容
	Visit      string // 文章阅读量
	Support    string // 文章点赞量
	CategoryId string // 文章分类id
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
	DeletedAt  string // 删除时间
}

//  blogArticleColumns holds the columns for table blog_article.
var blogArticleColumns = BlogArticleColumns{
	Id:         "id",
	Author:     "author",
	Cover:      "cover",
	Title:      "title",
	Summary:    "summary",
	Content:    "content",
	Visit:      "visit",
	Support:    "support",
	CategoryId: "category_id",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
}

// NewBlogArticleDao creates and returns a new DAO object for table data access.
func NewBlogArticleDao() *BlogArticleDao {
	return &BlogArticleDao{
		group:   "default",
		table:   "blog_article",
		columns: blogArticleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BlogArticleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BlogArticleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BlogArticleDao) Columns() BlogArticleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BlogArticleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BlogArticleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BlogArticleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
