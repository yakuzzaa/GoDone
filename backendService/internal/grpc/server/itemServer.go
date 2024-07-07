package server

import (
	"context"
	"github.com/yakuzzaa/GoDone/backendService/grpc/pkg/item_v1"
	"github.com/yakuzzaa/GoDone/backendService/internal/grpc/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ItemServer struct {
	item_v1.UnimplementedItemV1Server
	Service service.ItemServiceInterface
}

func NewItemServer(service service.ItemServiceInterface) *ItemServer {
	return &ItemServer{Service: service}
}

func (i *ItemServer) CreateItem(ctx context.Context, req *item_v1.CreateRequest) (*item_v1.CreateResponse, error) {
	createItem, err := i.Service.CreateItem(req.ListId, req.Info)
	if err != nil {
		return &item_v1.CreateResponse{}, status.Errorf(codes.Internal, "failed to create item: %v", err)
	}
	return &item_v1.CreateResponse{Id: createItem}, nil
}

func (i *ItemServer) ListItem(ctx context.Context, req *item_v1.ListRequest) (*item_v1.ListResponse, error) {
	items, err := i.Service.GetList(req.ListId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get item list: %v", err)
	}
	response := &item_v1.ListResponse{
		Items: []*item_v1.Item{},
	}
	for _, item := range *items {
		response.Items = append(response.Items, item)
	}

	return response, nil
}

func (i *ItemServer) GetItem(ctx context.Context, req *item_v1.GetRequest) (*item_v1.GetResponse, error) {
	item, err := i.Service.GetById(req.Id)
	if err != nil {
		return &item_v1.GetResponse{}, status.Errorf(codes.Internal, "failed to get item: %v", err)
	}
	response := &item_v1.GetResponse{
		Item: item,
	}
	return response, nil
}

func (i *ItemServer) UpdateItem(ctx context.Context, req *item_v1.UpdateRequest) (*emptypb.Empty, error) {
	err := i.Service.Update(req.Id, req.Info)
	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "failed to update item: %v", err)
	}
	return &emptypb.Empty{}, nil
}

func (i *ItemServer) DeleteItem(ctx context.Context, req *item_v1.DeleteRequest) (*emptypb.Empty, error) {
	err := i.Service.Delete(req.Id)
	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "failed to delete item: %v", err)
	}
	return &emptypb.Empty{}, nil
}
