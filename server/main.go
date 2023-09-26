package main

import (
	"context"
	"fmt"
	"net"
	"student/config"
	"student/controllers"
	pb "student/proto"
	"student/services"


	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func InitApp(client *mongo.Client) {
	ctx := context.TODO()
	Usercoll := config.GetCollection(client, "student", "User")
	service := services.InitUserSevice(Usercoll, ctx)
	controllers.UserService = service
}

func main() {
	client, err := config.Connectdatabase()
	if err != nil {
		fmt.Println(err)
	}
	defer client.Disconnect(context.TODO())
	fmt.Println("database connected")
	InitApp(client)
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &controllers.RPCServer{})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
	//defer client.Disconnect(context.TODO())
}
