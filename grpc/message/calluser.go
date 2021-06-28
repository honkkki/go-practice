// message服务使用user服务
package main

import (
	"context"
	"fmt"
	"log"

	messagePb "go-practice/grpc/message/pb"
	userPb "go-practice/grpc/user/pb"
	"google.golang.org/grpc"
)

func main() {
	// 连接用户服务
	conn1, err := grpc.Dial("127.0.0.1:9988", grpc.WithInsecure())
	// 连接消息服务
	conn2, err := grpc.Dial("127.0.0.1:9999", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("faild to connect server: %v", err)
	}
	defer conn1.Close()
	defer conn2.Close()

	uc := userPb.NewUserClient(conn1)
	mc := messagePb.NewGreeterClient(conn2)

	userInfo, err := uc.GetUserInfo(context.Background(), &userPb.UserReq{Id: 1})
	if userInfo == nil {
		log.Fatal("cant get user info, ", err)
	}
	msg, err := mc.SayHello(context.Background(), &messagePb.HelloRequest{Name: userInfo.Name})

	if err != nil {
		fmt.Printf("call fail: %v", err)
	}

	fmt.Printf("Greeting: %s!\n", msg.Message)
}
