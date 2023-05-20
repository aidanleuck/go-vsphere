package main

import (
	"context"
	"fmt"
	"vsphere_module/src/common"
	"vsphere_module/src/queue"
	"vsphere_module/src/routes"
	"vsphere_module/src/service"
	"vsphere_module/src/vimautomation"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	appConfig := common.GetConfiguration()

	mongoConnection, err := mongo.Connect(context.Background(), options.Client().ApplyURI(appConfig.DB.ConnectionString))
	if err != nil {
		panic(err)
	}
	lanierConnection, err := vimautomation.Connect(&appConfig.Lanier)
	if err != nil {
		panic(err)
	}

	service := &service.AgentService{
		VsphereClient: &vimautomation.VSphereClient{Client: lanierConnection},
		Config:        appConfig,
		Queue:         &queue.QueueModel{DB: mongoConnection},
		Logger:        common.InitLogger(),
	}

	_ = vimautomation.VSphereClient.Clone(*service.VsphereClient, "DC0", "/vm/DC0_H0_VM0")

	router := routes.SetupRouter(service)
	hostEndpoint := fmt.Sprintf("localhost:%d", service.Config.Port)
	service.Logger.Infow("Hosting agent at", "url", hostEndpoint)
	routerErr := router.Run(hostEndpoint)
	if routerErr != nil {
		panic(fmt.Errorf("fatal error starting router: %w", routerErr))
	}

}
