package services

import (
	"context"
	"fmt"
	"strings"
	interfaces "student/Interfaces"
	"student/models"

	//"student/models"

	"go.mongodb.org/mongo-driver/bson"
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

func (u*UserService)FindUser(user int)([]*models.User,error){
match := bson.D{{
	Key:   "UserId",
	Value: user,
}}
res,err:=u.UserService.Find(u.ctx,match)
if err!=nil{
	return nil,err
}else{
 var user_details[]*models.User
 for res.Next(u.ctx){
	details := &models.User{}
	err:= res.Decode(details)
	if err!= nil{
		return nil,err
	}
	user_details = append(user_details,details)
 }

return user_details,nil

}

}