// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: confRpc.proto

package server

import (
	"context"

	"zero-online-conf/confRpc/confRpc"
	"zero-online-conf/confRpc/internal/logic"
	"zero-online-conf/confRpc/internal/svc"
)

type ConfRpcServer struct {
	svcCtx *svc.ServiceContext
	confRpc.UnimplementedConfRpcServer
}

func NewConfRpcServer(svcCtx *svc.ServiceContext) *ConfRpcServer {
	return &ConfRpcServer{
		svcCtx: svcCtx,
	}
}

func (s *ConfRpcServer) Ping(ctx context.Context, in *confRpc.Request) (*confRpc.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
