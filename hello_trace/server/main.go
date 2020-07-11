package main

import (
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/net/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"hello_trace/proto/hello"
	"log"
	"net"
	"net/http"
)

const (
	Address = "127.0.0.1:8998"
	NetWork = "tcp"
)

type helloService struct{}

func (h helloService) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloResponse, error) {
	res := &hello.HelloResponse{}
	res.Message = fmt.Sprintf("Hello %s.", in.Name)
	return res, nil
}

var HelloService = helloService{}

func main() {
	listener, err := net.Listen(NetWork, Address)
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()

	hello.RegisterHelloServiceServer(server, HelloService)

	//todo
	go startTrace()

	log.Println("Listen on" + Address)
	grpclog.Info("Listen on" + Address)
	server.Serve(listener)
}

func startTrace() {
	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true, true
	}

	go http.ListenAndServe(":8998", nil)
	log.Println("Trace listen on 8998")
	grpclog.Info("Trace listen on 8998")
}
