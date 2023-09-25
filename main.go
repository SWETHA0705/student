package main

import (
	//"fmt"
	"fmt"
	"student/config"
)

func main(){
	client,_ := config.Connectdatabase()
	collection :=config.GetCollection(client,"students","details")
	fmt.Println("database connected,",collection)
	



}