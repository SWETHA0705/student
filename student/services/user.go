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

type UserService struct {
	UserService *mongo.Collection
	ctx         context.Context
}

func InitUserSevice(collection *mongo.Collection, ctx context.Context) interfaces.IUser {
	return &UserService{collection, ctx}

}

func (u *UserService) CreateUser(user *models.User) (string, error) {
	user.Email = strings.ToLower(user.Email)
	_, err := u.UserService.InsertOne(u.ctx, user)
	if err != nil {
		fmt.Println(err)
		return "nil", err
	}
	return "success", nil
}

func (u *UserService) FindUser(user int) (*models.User, error) {
	filter := bson.D{{Key: "UserId", Value: user}}
	var find *models.User
	res := u.UserService.FindOne(u.ctx, filter)
	err := res.Decode(&find)
	if err != nil {
		return nil, err
	}
	return find, nil
}

func (u *UserService) UpdateUser(user int, role string) (string, error) {
	filter := bson.D{{
		Key:   "UserId",
		Value: user,
	}}
	fv := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "Role", Value: role},
		}},
	}
	//var update *models.User
	res, err := u.UserService.UpdateOne(u.ctx, filter, fv)
	if res.MatchedCount == 0 {
		// No documents matched the filter criteria, you can return an error or handle it accordingly.
		return "nil", fmt.Errorf("No user found for UserId: %d", user)
	}

	fmt.Println(res)
	if err != nil {
		return "nil", err
	}

	return "sucess", nil
}
