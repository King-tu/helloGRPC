package main

import (
	"context"
	pb "github.com/king-tu/helloGRPC/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type HelloService struct{}

func (*HelloService) Hello(ctx context.Context, msg *pb.StringMessage) (*pb.StringMessage, error) {
	log.Println("client request message: ", msg)
	reply := &pb.StringMessage{
		Value: "hello " + msg.Value,
	}
	return reply, nil
}

func main() {
	address := ":1234"
	//获取grpc服务端对象
	grpcServer := grpc.NewServer()
	//注册grpc服务
	pb.RegisterHelloServiceServer(grpcServer, new(HelloService))

	//设置服务端监听
	lis, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	//在指定端口上提供grpc服务
	log.Println("gRPC Server listen on: ", address)
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
