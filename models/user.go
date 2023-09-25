package models

type User struct {
	UserId int `json:"UserId" bson:"UserId"`
	Email string `json:"Email" bson:"Email"`
	PhoneNum int `json:"PhoneNum" bson:"PhoneNum"`
	Role string `json:"Role" bson:"Role"`
}
