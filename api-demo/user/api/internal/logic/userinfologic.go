package logic

import (
	"context"

	"go-zero-demo/api-demo/user/api/internal/svc"
	"go-zero-demo/api-demo/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line
	return &types.UserInfoResponse{Code: 200, Data: types.UserInfo{
		Id:       1,
		UserName: "kevin",
		Addr:     "广东深圳",
	}, Msg: "success"}, nil
}