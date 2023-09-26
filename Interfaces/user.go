package interfaces

import (
	"student/models"

	//"go.mongodb.org/mongo-driver/mongo"
)

// "os/user"
// "student/models"

//import "go.mongodb.org/mongo-driver/mongo"

type IUser interface{
	CreateUser(*models.User)(string,error)
	FindUser(int)([]*models.User,error)
}