// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"github.com/SliverFlow/minidouyin/app/user/cmd/rpc/internal/logic"
	"github.com/SliverFlow/minidouyin/app/user/cmd/rpc/internal/svc"
	"github.com/SliverFlow/minidouyin/app/user/cmd/rpc/pb"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

// 创建用户
func (s *UserServer) Create(ctx context.Context, in *pb.CreateReq) (*pb.CreateReply, error) {
	l := logic.NewCreateLogic(ctx, s.svcCtx)
	return l.Create(in)
}
