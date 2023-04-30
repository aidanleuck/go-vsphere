package routes

import (
	"net/http"
	"vsphere_module/src/common"
	"vsphere_module/src/queue"

	"github.com/gin-gonic/gin"
)

func GetJobs(serv *common.AgentService) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := serv.Queue.GetAllJobs()
		c.IndentedJSON(http.StatusOK, result)
	}
}

func AddJob(serv *common.AgentService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var jobToAdd *queue.Job = new(queue.Job)
		c.BindJSON(jobToAdd)
		serv.Queue.AddJob(jobToAdd)
	}
}

func FinishJob(serv *common.AgentService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var jobID string
		c.BindJSON(jobID)

		discoveredJob, err := serv.Queue.GetJob(jobID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}

		serv.Queue.FinishJob(discoveredJob)
	}
}
