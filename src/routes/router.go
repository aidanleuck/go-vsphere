package routes

import (
	"vsphere_module/src/common"

	"github.com/gin-gonic/gin"
)

func SetupRouter(srv *common.AgentService) *gin.Engine {
	router := gin.Default()
	SetupGraphQLRouter(router, srv)
	return router
}
