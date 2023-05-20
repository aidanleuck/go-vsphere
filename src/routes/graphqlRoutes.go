package routes

import (
	"vsphere_module/src/service"

	"github.com/gin-gonic/gin"
)

func SetupGraphQLRouter(router *gin.Engine, service *service.AgentService) {
	router.GET("/", playgroundHandler())
	router.POST("/query", graphqlHandler(service))
}
