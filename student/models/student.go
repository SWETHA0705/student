package models

type Students struct{
	Name  string`json:"Name" bson:"Name"`
	RollNo int `json:"RollNo" bson:"RollNo"`
	Department string `json:"Department" bson:"Department"`
	Address []Address `json:"Address" bson:"Address"`
}

type Address struct{
 DoorNo int `json:"DoorNo" bson:"DoorNo"`
 street string `json:"street" bson:"street"`
}