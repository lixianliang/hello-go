package server

import (
	"context"
	"log"
	"net"

	//"gokit/pkg/model"
	"github.com/lixianliang/hello-go/go-kit/grpc-03/pkg/model"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// User ...
type User struct {
	Name    string
	Title   string
	Company string
}

type UManagerServiceServer struct {
	Users []*User
}

func (server *UManagerServiceServer) GetUser(ctx context.Context, req *model.GetUserRequest) (res *model.GetUserResponse, err error) {
	for _, u := range server.Users {
		if u.Name == req.Name {
			return &model.GetUserResponse{
				Name:    u.Name,
				Title:   u.Title,
				Company: u.Company,
			}, nil
		}
	}
	return nil, model.ErrUserNotFound
}

func (server *UManagerServiceServer) SetTitle(ctx context.Context, req *model.SetTitleRequest) (res *model.Empty, err error) {
	//return &model.Empty{}, model.UM.SetTitle(req.Name, req.Title)
	for _, u := range server.Users {
		if u.Name == req.Name {
			u.Title = req.Title
			return &model.Empty{}, nil
		}
	}
	return &model.Empty{}, model.ErrUserNotFound
}

func (server *UManagerServiceServer) Dispatch(ctx context.Context, req *model.DispatchRequest) (res *model.Empty, err error) {
	//return &model.Empty{}, model.UM.Dispatch(req.Name, req.Company)
	for _, u := range server.Users {
		if u.Name == req.Name {
			u.Company = req.Company
			return &model.Empty{}, nil
		}
	}
	return &model.Empty{}, model.ErrUserNotFound
}

func NewServer() {
	log.Println("server: 启动")
	lis, err := net.Listen("tcp", model.ServerAddr)
	if err != nil {
		panic(err)
	}
	rpcServer := grpc.NewServer()
	log.Println("server: 注册服务")
	//model.RegisterUserManagerServiceServer(rpcServer, &UManagerServiceServer{})
	model.RegisterUserManagerServiceServer(rpcServer, uManagerServiceServer)
	reflection.Register(rpcServer)
	log.Println("server: 等待连接")
	if err := rpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
