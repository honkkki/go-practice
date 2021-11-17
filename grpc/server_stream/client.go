package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "go-practice/grpc/server_stream/pb"
	"google.golang.org/grpc"
)

func SayList(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayHelloList(context.Background(), r)
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("resp: %v\n", resp)
		time.Sleep(time.Second)
	}

	return nil
}

func main() {
	// 连接服务器
	conn, err := grpc.Dial("127.0.0.1:9999", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	// 调用服务端的SayHello
	req := &pb.HelloRequest{Name: "karina"}
	err = SayList(c, req)
	if err != nil {
		fmt.Printf("could not greet: %v", err)
	}
}
