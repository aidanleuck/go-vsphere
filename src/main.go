package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type agentService struct {
	config *appConfiguration
	queue  Queue
}

func main() {

	appConfig := GetConfiguration()

	mongoConnection, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(appConfig.DB.ConnectionString))
	service := &agentService{
		queue: QueueModel{DB: mongoConnection},
	}

	service.queue.GetAll()

	router := gin.Default()
	hostEndpoint := fmt.Sprintf("localhost:%d", appConfig.Port)
	println("Hosting agent at", hostEndpoint)

	router.GET("/job", GetJobs)
	router.POST("/job", AddJob)

	routerErr := router.Run(hostEndpoint)
	if routerErr != nil {
		panic(fmt.Errorf("fatal error starting router: %w", err))
	}

}
