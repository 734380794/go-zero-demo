package logic

import (
	"context"
	"go-zero-demo/common/jwts"

	"go-zero-demo/api-demo/user/api_jwt/internal/svc"
	"go-zero-demo/api-demo/user/api_jwt/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp string, err error) {
	// todo: add your logic here and delete this line

	auth := l.svcCtx.Config.Auth
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		UserID:   1,
		Username: "kevin",
		Role:     1}, auth.AccessSecret, auth.AccessExpire)
	if err != nil {
		return "", nil
	}
	return token, nil
}
