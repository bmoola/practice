package main

import (
	"log"
	"net"

	"github.com/practice/example1/apis/protos"
	"github.com/practice/example1/services/service"
	"google.golang.org/grpc"
)

func main(){
	lis, err := net.Listen("tcp", ":50051")
	if err != nil{
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()

	s := service.NewUserServiceServer()

	protos.RegisterGreetingServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil{
		log.Fatalf("failed to serve: %v", err)
	}
}