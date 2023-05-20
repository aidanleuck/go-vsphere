package routes

import (
	"vsphere_module/src/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(srv *service.AgentService) *gin.Engine {
	router := gin.Default()
	SetupGraphQLRouter(router, srv)
	return router
}
