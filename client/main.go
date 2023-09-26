package main

import (
	"context"
	"net/http"
	pb "student/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	r := gin.Default()

	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewUserClient(conn)
	r.POST("/create", func(c *gin.Context) {
		var request pb.UserRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		res, err := client.CreateUser(context.TODO(), &request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": res.Success})
	})
	r.Run(":2000")

	r.POST("/find",func(c*gin.Context){
		
	})
}
