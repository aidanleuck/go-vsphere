package main

import (
	"context"
	"fmt"
	"vsphere_module/src/common"
	"vsphere_module/src/queue"
	"vsphere_module/src/routes"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	appConfig := common.GetConfiguration()

	mongoConnection, err := mongo.Connect(context.Background(), options.Client().ApplyURI(appConfig.DB.ConnectionString))

	if err != nil {
		panic(err)
	}

	service := &common.AgentService{
		Config: appConfig,
		Queue:  &queue.QueueModel{DB: mongoConnection},
	}

	router := routes.SetupRouter(service)
	hostEndpoint := fmt.Sprintf("localhost:%d", service.Config.Port)
	appConfig.Logger.Infow("Hosting agent at", "url", hostEndpoint)
	routerErr := router.Run(hostEndpoint)
	if routerErr != nil {
		panic(fmt.Errorf("fatal error starting router: %w", routerErr))
	}

}
