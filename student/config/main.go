package config

import (
	"context"
	"fmt"
	//"errors"
	"log"
	constants "student/Constants"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Connectdatabase()(*mongo.Client,error){
	ctx,_ := context.WithTimeout(context.Background(),10*time.Second)
	mongoconnection:= options.Client().ApplyURI(constants.Connectionstring)
	mongoclient, err := mongo.Connect(ctx,mongoconnection)
	if err!= nil{
		log.Fatal(err.Error())
		return nil,err
	}
	if err:= mongoclient.Ping(ctx,readpref.Primary()); err!= nil{
       return nil, err
	}
	return mongoclient,nil

}

func GetCollection(client *mongo.Client,dbName string,collectionName string)*mongo.Collection{
	client,err := Connectdatabase()
	if err!= nil{
		fmt.Println(err)
		panic(err)
	}
	collection:= client.Database(dbName).Collection(collectionName)
	return collection
}