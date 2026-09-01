package rpc

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	"github.com/okm321/mahking/go/internal/application"
	groupv1 "github.com/okm321/mahking/go/internal/presentation/rpc/gen/mahking/group/v1"
	"github.com/okm321/mahking/go/internal/presentation/rpc/gen/mahking/group/v1/groupv1connect"
	pkgerror "github.com/okm321/mahking/go/pkg/error"
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

func (s *GroupServer) GetGroup(ctx context.Context, req *connect.Request[groupv1.GetGroupRequest]) (*connect.Response[groupv1.GetGroupResponse], error) {
	group, err := s.usecase.Get(ctx, req.Msg.Uid)
	if err != nil {
		if errors.Is(err, pkgerror.ErrNotFound) {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&groupv1.GetGroupResponse{
		Group: &groupv1.Group{
			Id:   group.ID,
			Uid:  group.UID,
			Name: group.Name,
		},
	}), nil
}
