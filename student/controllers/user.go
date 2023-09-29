package controllers

import (
	"context"
	//	"strings"
	interfaces "student/Interfaces"
	"student/models"
	pb "student/proto"
	//"golang.org/x/text/message"
)

type RPCServer struct {
	pb.UnimplementedUserServer
}

var (
	UserService interfaces.IUser
)

func (s *RPCServer) CreateUser(ctx context.Context, u *pb.UserRequest) (*pb.UserResponse, error) {
	user := &models.User{
		UserId:   u.UserId,
		Email:    u.Email,
		PhoneNum: u.PhoneNum,
		Role:     u.Role,
	}
	_, err := UserService.CreateUser(user)
	if err != nil {
		return &pb.UserResponse{
			Success: "failed",
		}, err
	}
	return &pb.UserResponse{
		Success: "inserted",
	}, nil
}

func (s *RPCServer) FindUser(ctx context.Context, f *pb.FindRequest) (*pb.FindResponse, error) {
	// User:= &models.User{
	// 	UserId:   f.UserId,

	// }
	res, err := UserService.FindUser(int(f.UserId))
	if err != nil {
		return &pb.FindResponse{}, err
	}
	return &pb.FindResponse{
		UserId: res.UserId,
		Email:  res.Email,
	}, nil
}

func (s *RPCServer) UpdateUser(ctx context.Context, u *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	user := &models.User{
		UserId: u.UserId,
		Role:   u.Role,
	}
	_, err := UserService.UpdateUser(int(user.UserId), user.Role)
	if err != nil {
		return &pb.UpdateResponse{}, err

	}
	return &pb.UpdateResponse{
		Message: "updated",
	}, err
}

