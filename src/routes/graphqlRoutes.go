package routes

import (
	"vsphere_module/src/common"

	"github.com/gin-gonic/gin"
)

func SetupGraphQLRouter(router *gin.Engine, service *common.AgentService) {
	router.GET("/", playgroundHandler())
	router.POST("/query", graphqlHandler(service))
}
