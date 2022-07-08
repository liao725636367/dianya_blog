package service

import (
	"blog/internal/consts"
	"blog/internal/model/entity"
	"context"
	"fmt"
	"time"

	"github.com/rs/xid"
)

type IToken interface {
	GetUserByToken(ctx context.Context, token string) (*entity.BlogUser, error)
	RefreshToken(ctx context.Context, token string, userid uint) error
	CreateToken(ctx context.Context, userid uint) (token string, err error)
}
type (
	sToken    struct{}
	tokenInfo struct {
		Userid uint  `json:"userid"`
		Expire int64 `json:"expire"`
	}
)

var (
	insToken = sToken{}
)

func Token() IToken {
	return &insToken
}

// GetUserByToken 根据token 获取用户信息
func (s *sToken) GetUserByToken(ctx context.Context, token string) (*entity.BlogUser, error) {
	tokenDataVar, err := cacheClient.Get(ctx, token)
	if err != nil {
		return nil, err
	}
	var tokenData = new(tokenInfo)
	err = tokenDataVar.Struct(tokenData)
	if err != nil {
		return nil, err
	}
	//token 过期 续签 token
	if time.Now().After(time.Unix(tokenData.Expire, 0)) {
		s.RefreshToken(ctx, token, tokenData.Userid)
	}

	return User().GetUserInfo(ctx, tokenData.Userid)
}

// RefreshToken 刷新token过期时间
func (s *sToken) RefreshToken(ctx context.Context, token string, userid uint) error {
	tokenData := &tokenInfo{
		Userid: userid,
		Expire: time.Now().Add(time.Second * consts.TokenExpire).Unix(),
	}
	return cacheClient.Set(ctx, token, tokenData, time.Second*consts.TokenExpire*2)
}

// CreateToken 创建token
func (s *sToken) CreateToken(ctx context.Context, userid uint) (token string, err error) {
	token = fmt.Sprintf("%s_%d", xid.New(), userid)
	tokenData := &tokenInfo{
		Userid: userid,
		Expire: time.Now().Add(time.Second * consts.TokenExpire).Unix(),
	}
	return token, cacheClient.Set(ctx, token, tokenData, time.Second*consts.TokenExpire)
}
