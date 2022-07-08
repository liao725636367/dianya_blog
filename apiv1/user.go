package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type UserLoginReq struct {
	g.Meta    `path:"/user/login" method:"post" tags:"用户" summary:"用户登录"`
	Username  string `json:"username" v:"required|length:1,20"`
	Password  string `json:"password" v:"required|password"`
	Password1 string `json:"password1" v:"required|password|same:Password" `
}
type UserLoginRes struct {
	Token string `json:"token"`
}

type UserRow struct {
	Id        uint        `json:"id"        ` // 主键id
	Username  string      `json:"username"  ` // 用户名
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
	RoleType  int         `json:"roleType"  ` // 用户角色 0普通用户 1管理员
}

type UserInfoReq struct {
	g.Meta `path:"/user" method:"get" tags:"用户" summary:"用户信息"`
}
type UserInfoRes UserRow
