package routes

import (
	"vsphere_module/src/common"

	"github.com/gin-gonic/gin"
)

func setupQueueRouter(router *gin.Engine, service *common.AgentService) {
	router.GET("/job", GetJobs(service))
	router.POST("/job", AddJob(service))
	router.POST("/job/finish")
}
