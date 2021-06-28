package main

import (
	"context"
	"fmt"

	pb "go-practice/grpc/user/pb"
	"google.golang.org/grpc"
)

func main() {
	// 连接服务器
	conn, err := grpc.Dial("127.0.0.1:9988", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewUserClient(conn)

	r, err := c.GetUserInfo(context.Background(), &pb.UserReq{Id: 1})
	if err != nil {
		fmt.Printf("get user info fail: %v", err)
	}
	fmt.Printf("user_id: %v user_name: %v \n", r.Id, r.Name)
}
