package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type vm struct {
	Name string
}

func GetJobs(c *gin.Context) {
	result := QueueModel.GetAll(QueueModel{})
	c.IndentedJSON(http.StatusOK, result)
}

func AddJob(c *gin.Context) {
	var jobToAdd *job = new(job)
	c.BindJSON(jobToAdd)

	//Queue.AddNewJob(jobToAdd)
}
