package controller

import (
	"blog/apiv1"
	"blog/internal/service"
	"blog/utility/utils"
	"context"
)

type cUser struct {
}

var User = cUser{}

//用户登录
func (c *cUser) Login(ctx context.Context, in *apiv1.UserLoginReq) (out *apiv1.UserLoginRes, err error) {
	token, err := service.User().Login(ctx, in)
	if err != nil {
		return nil, err
	}
	out = &apiv1.UserLoginRes{Token: token}
	return out, nil
}
func (c *cUser) UserInfo(ctx context.Context, in *apiv1.UserInfoReq) (out *apiv1.UserInfoRes, err error) {
	userid := service.User().CtxUserid(ctx)
	if userid == 0 {
		return nil, utils.CodeErr(401, "请登录")
	}
	row, err := service.User().GetUserInfo(ctx, userid)
	if err != nil {
		return nil, err
	}
	out = &apiv1.UserInfoRes{
		Id:        row.Id,
		Username:  row.Username,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		RoleType:  row.RoleType,
	}
	return out, nil
}
