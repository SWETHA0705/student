package services

import (
	"context"
	"fmt"
	"strings"
	interfaces "student/Interfaces"
	"student/models"

	//"student/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct{
	UserService *mongo.Collection
	ctx context.Context

}

func InitUserSevice (collection *mongo.Collection, ctx context.Context)interfaces.IUser{
	return &UserService{collection,ctx}
	
}

func (u * UserService)CreateUser(user *models.User)(string,error){
	user.Email = strings.ToLower(user.Email)
	_,err:= u.UserService.InsertOne(u.ctx,user)
    if err!= nil{
		fmt.Println(err)
		return "nil",err
	}
	return "success",nil
}