package main

import (
	"context"
	pb "github.com/king-tu/helloGRPC/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	//和grpc服务建立连接
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &pb.StringMessage{Value: "lisi"})
	if err != nil {
		panic(err)
	}
	log.Println("reply: ", reply)
}
