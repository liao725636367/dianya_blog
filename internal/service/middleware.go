/*
* @desc:中间件处理
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/3/17 9:11
 */

package service

import (
	"blog/internal/consts"
	"fmt"
	"net/http"

	"github.com/gogf/gf/v2/util/gvalid"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
)

type IMiddleware interface {
	MiddlewareCORS(r *ghttp.Request)
	Ctx(r *ghttp.Request)
	MiddlewareHandlerResponse(r *ghttp.Request)
}

type sMiddleware struct{}

var middleService = sMiddleware{}

func Middleware() IMiddleware {
	return &middleService
}

//跨域访问支持
func (s *sMiddleware) MiddlewareCORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	// you can set options
	//corsOptions.AllowDomain = []string{"goframe.org", "baidu.com"}
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}

//管理员上下文写入
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	token := r.GetHeader(consts.UserToken)
	if token == "" {
		r.Middleware.Next()
		return
	}
	ctx := r.Context()
	//获取用户信息并将用户id和是否管理员写到上下文
	userData, err := Token().GetUserByToken(ctx, token)
	if err != nil {
		glog.Error(ctx, fmt.Sprintf("token %s 解析错误：%v", token, err))
		r.Middleware.Next()
		return
	}
	r.SetCtxVar("userid", userData.Id)
	r.SetCtxVar("tole_type", userData.RoleType)
	r.SetCtxVar("username", userData.Username)
	r.Middleware.Next()
}

//管理员权限验证
func (s *sMiddleware) Auth(r *ghttp.Request) {
	if User().IsCtxAdmin(r.Context()) {
		r.Middleware.Next()
	} else {
		_ = r.Response.WriteJsonExit(http.StatusForbidden)
	}
}

type DefaultHandlerResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// MiddlewareHandlerResponse is the default middleware handling handler response object and its error.
func (s *sMiddleware) MiddlewareHandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		err         error
		res         interface{}
		ctx         = r.Context()
		internalErr error
	)
	res, err = r.GetHandlerResponse()
	if err != nil {
		code := gerror.Code(err)
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		msg := err.Error()
		switch realErr := err.(type) {
		case gvalid.Error:
			msg = realErr.FirstError().Error()
		}
		internalErr = r.Response.WriteJson(DefaultHandlerResponse{
			Code:    code.Code(),
			Message: msg,
			Data:    nil,
		})
		if internalErr != nil {
			glog.Error(ctx, internalErr)
		}
		return
	}
	internalErr = r.Response.WriteJson(DefaultHandlerResponse{
		Code:    gcode.CodeOK.Code(),
		Message: "ok",
		Data:    res,
	})
	if internalErr != nil {
		glog.Error(ctx, internalErr)
	}
}
