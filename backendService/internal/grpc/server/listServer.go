package server

import (
	"context"

	"github.com/yakuzzaa/GoDone/backendService/grpc/pkg/list_v1"
	"github.com/yakuzzaa/GoDone/backendService/internal/grpc/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ListServer struct {
	list_v1.UnimplementedListV1Server
	Service service.ListServiceInterface
}

func NewListServer(service service.ListServiceInterface) *ListServer {
	return &ListServer{Service: service}
}

func (l *ListServer) CreateList(ctx context.Context, req *list_v1.CreateRequest) (*list_v1.CreateResponse, error) {
	createList, err := l.Service.CreateList(req.Info, req.UserId)
	if err != nil {
		return &list_v1.CreateResponse{}, status.Errorf(codes.Internal, "failed to create list: %v", err)
	}
	return &list_v1.CreateResponse{Id: createList}, nil
}

func (l *ListServer) List(ctx context.Context, req *list_v1.ListRequest) (*list_v1.ListResponse, error) {
	lists, err := l.Service.GetList(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get list: %v", err)
	}
	response := &list_v1.ListResponse{
		Lists: []*list_v1.List{},
	}
	for _, list := range *lists {
		response.Lists = append(response.Lists, list)
	}

	return response, nil
}

func (l *ListServer) GetDetail(ctx context.Context, req *list_v1.DetailRequest) (*list_v1.DetailResponse, error) {
	list, items, err := l.Service.GetListById(req.Id, req.UserId)
	if err != nil {
		return &list_v1.DetailResponse{}, status.Errorf(codes.Internal, "failed to get detail: %v", err)
	}
	response := &list_v1.DetailResponse{
		List: &list_v1.ListWithItems{
			List:  list,
			Items: items,
		},
	}
	return response, nil
}

func (l *ListServer) UpdateList(ctx context.Context, req *list_v1.UpdateRequest) (*emptypb.Empty, error) {
	err := l.Service.Update(req.Id, req.UserId, req.Info)
	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "failed to update list: %v", err)
	}
	return &emptypb.Empty{}, nil
}

func (l *ListServer) DeleteList(ctx context.Context, req *list_v1.DeleteRequest) (*emptypb.Empty, error) {
	err := l.Service.Delete(ctx, req.Id, req.UserId)
	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "failed to delete list: %v", err)
	}
	return &emptypb.Empty{}, nil
}
