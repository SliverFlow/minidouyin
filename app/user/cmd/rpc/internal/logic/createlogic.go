package logic

import (
	"context"

	"github.com/SliverFlow/minidouyin/app/user/cmd/rpc/internal/svc"
	"github.com/SliverFlow/minidouyin/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建用户
func (l *CreateLogic) Create(in *pb.CreateReq) (*pb.CreateReply, error) {
	// todo: add your logic here and delete this line

	return &pb.CreateReply{}, nil
}
