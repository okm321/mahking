package rpc

import (
	"context"

	"connectrpc.com/connect"
	"github.com/okm321/mahking/go/internal/application"
	groupv1 "github.com/okm321/mahking/go/internal/presentation/rpc/gen/mahking/group/v1"
	"github.com/okm321/mahking/go/internal/presentation/rpc/gen/mahking/group/v1/groupv1connect"
	pkgtrace "github.com/okm321/mahking/go/pkg/trace"
)

type GroupServer struct {
	groupv1connect.UnimplementedGroupServiceHandler
	usecase *application.GroupUsecase
}

func NewGroupServer(usecase *application.GroupUsecase) *GroupServer {
	return &GroupServer{
		usecase: usecase,
	}
}

func (s *GroupServer) GetGroup(ctx context.Context, req *connect.Request[groupv1.GetGroupRequest]) (_ *connect.Response[groupv1.GetGroupResponse], err error) {
	ctx = pkgtrace.StartSpan(ctx, "GroupServer.GetGroup")
	defer func() { pkgtrace.EndSpan(ctx, err) }()

	group, err := s.usecase.Get(ctx, req.Msg.Uid)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&groupv1.GetGroupResponse{
		Group: toProtoGroup(group),
	}), nil
}

func (s *GroupServer) CreateGroup(ctx context.Context, req *connect.Request[groupv1.CreateGroupRequest]) (_ *connect.Response[groupv1.CreateGroupResponse], err error) {
	ctx = pkgtrace.StartSpan(ctx, "GroupServer.CreateGroup")
	defer func() { pkgtrace.EndSpan(ctx, err) }()

	group, err := s.usecase.Create(ctx, toCreateGroupInput(req.Msg))
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&groupv1.CreateGroupResponse{
		Group: toProtoGroup(group),
	}), nil
}
