package logic

import (
	"context"

	"github.com/SliverFlow/minidouyin/app/user/cmd/api/internal/svc"
	"github.com/SliverFlow/minidouyin/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.UserCreateReq) (resp *types.UserCreateReply, err error) {
	// todo: add your logic here and delete this line

	return
}
