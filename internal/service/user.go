package service

import (
	"blog/apiv1"
	"blog/internal/consts"
	"blog/internal/model/entity"
	"blog/internal/service/internal/dao"
	"blog/internal/service/internal/do"
	"blog/utility/utils"
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/database/gdb"
)

type IUser interface {
	GetUserInfo(ctx context.Context, userid uint) (*entity.BlogUser, error)
	IsCtxAdmin(ctx context.Context) bool
	Login(ctx context.Context, in *apiv1.UserLoginReq) (string, error)
	CtxUserid(ctx context.Context) uint
}
type (
	sUser struct{}
)

var (
	insUser = sUser{}
)

func User() IUser {
	return &insUser
}

// GetUserInfo 获取用户信息
func (u *sUser) GetUserInfo(ctx context.Context, userid uint) (*entity.BlogUser, error) {
	var user *entity.BlogUser
	err := dao.BlogUser.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Hour,
		Name:     fmt.Sprintf(consts.RedisUserRow, userid),
		Force:    false,
	}).Where(do.BlogUser{
		Id: userid,
	}).Scan(&user)
	if err != nil {
		return nil, err
	}
	if user == nil {

		return nil, utils.CodeErr(1, "用户不存在")
	}
	return user, nil
}

// IsCtxAdmin 是否后台管理员
func (u *sUser) IsCtxAdmin(ctx context.Context) bool {
	v := ctx.Value("tole_type")
	roleType, ok := v.(int)
	if ok && roleType == consts.RoleAdmin {
		return true
	}
	return false
}

// CtxUserid 上下文获取用户id
func (u *sUser) CtxUserid(ctx context.Context) uint {
	v := ctx.Value("userid")
	return gconv.Uint(v)

}

// Login 用户登录
func (u *sUser) Login(ctx context.Context, in *apiv1.UserLoginReq) (string, error) {
	//结构体校验
	err1 := g.Validator().Data(in).Run(ctx)
	if err1 != nil {
		return "", err1.FirstError()
	}
	var user *entity.BlogUser
	//获取用户
	err := dao.BlogUser.Ctx(ctx).Where("username=? and status =1", in.Username).Scan(&user)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", utils.CodeErr(1, "用户不存在")
	}
	encodePassword := utils.EncodePassword(in.Password, user.Salt)
	if encodePassword != user.Password {
		return "", utils.CodeErr(1, "密码错误："+encodePassword)
	}
	return Token().CreateToken(ctx, user.Id)
}
