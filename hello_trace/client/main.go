package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"hello_trace/proto/hello"
)

const (
	Address = "127.0.0.1:8998"
)

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()

	cli := hello.NewHelloServiceClient(conn)

	req := &hello.HelloRequest{Name: "gRPC"}
	res, err := cli.SayHello(context.Background(), req)
	if err != nil {
		grpclog.Fatalln(err)
	}

	fmt.Println(res.Message)
}
