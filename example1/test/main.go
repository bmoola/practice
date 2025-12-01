package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/practice/example1/apis/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main(){
conn, err := grpc.Dial(
		"localhost:50051",                  // server address
		grpc.WithTransportCredentials(insecure.NewCredentials()), // no TLS for local dev
	)
	if err != nil {
		log.Fatalf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// 2) Create a client from the generated code.
	client := protos.NewGreetingClient(conn)

	// 3) Create a context with timeout for the RPC call.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.GetGreeting(ctx, &protos.UserInput{
		Firstname: "bhagi",
		Lastname: "Moola",
	})
	if err != nil{
		log.Fatalf("error occured")
	}

	fmt.Println(resp.Greeting)
}
