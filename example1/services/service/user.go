package service

import (
	"context"
	"fmt"

	"github.com/practice/example1/apis/protos"
)

type UserServiceServer struct{
	protos.UnimplementedGreetingServer
}

func NewUserServiceServer() UserServiceServer{
	return UserServiceServer{}
}

func (u *UserServiceServer) GetGreeting(ctx context.Context, c *protos.UserInput) (*protos.UserOutput, error){
	if c.Firstname != "" && c.Lastname != ""{
		return &protos.UserOutput{
			Greeting: fmt.Sprintf("Hello: %s %s", c.Firstname, c.Lastname),
		}, nil
	}

	return nil, fmt.Errorf("first name and last name are null")
}