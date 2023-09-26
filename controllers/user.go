package controllers

import (
	"context"
	interfaces "student/Interfaces"
	"student/models"
	pb "student/proto"

	//"golang.org/x/text/message"
)

type RPCServer struct{
 pb.UnimplementedUserServer
}

var(
	UserService interfaces.IUser
)

func(s* RPCServer)CreateUser(ctx context.Context,u *pb.UserRequest)(*pb.UserResponse,error){
	user:= &models.User{
		UserId:   u.UserId,
		Email:    u.Email,
		PhoneNum: u.PhoneNum,
		Role:     u.Role,
	}
	_,err:= UserService.CreateUser(user)
	if err!= nil{
		return &pb.UserResponse{
			Success: "failed",
		},err
	}
	return &pb.UserResponse{
		Success: "inserted",
	},nil
}

func (s*RPCServer)FindUser(ctx context.Context, f *pb.FindRequest)(*pb.FindResponse,error){
	
}