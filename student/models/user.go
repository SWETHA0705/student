package models

type User struct {
	UserId int32 `json:"UserId" bson:"UserId"`
	Email string `json:"Email" bson:"Email"`
	PhoneNum int32 `json:"PhoneNum" bson:"PhoneNum"`
	Role string `json:"Role" bson:"Role"`
}
