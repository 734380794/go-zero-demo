package logic

import (
	"context"
	"fmt"

	"go-zero-demo/api-demo/user/api/internal/svc"
	"go-zero-demo/api-demo/user/api/internal/types"

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

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	fmt.Println(req.UserName, req.Password)
	return &types.Response{Code: 200, Data: "xx.xx.xx", Msg: "success"}, nil
}
