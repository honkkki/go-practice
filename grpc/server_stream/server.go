package main

import (
	"fmt"
	"net"

	pb "go-practice/grpc/server_stream/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) SayHelloList(r *pb.HelloRequest, stream pb.Greeter_SayHelloListServer) error {
	for n := 0; n < 6; n++ {
		_ = stream.Send(&pb.HelloReply{Message: "hello " + r.Name})
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()                  // 创建gRPC服务器
	pb.RegisterGreeterServer(s, &server{}) // 在gRPC服务端注册服务

	reflection.Register(s) //在给定的gRPC服务器上注册服务器反射服务
	// Serve方法在lis上接受传入连接，为每个连接创建一个ServerTransport和server的goroutine。
	// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应它们。
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v\n", err)
		return
	}
}
