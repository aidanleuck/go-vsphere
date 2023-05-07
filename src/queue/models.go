package queue

import (
	"vsphere_module/src/graph/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type QueueModel struct {
	DB *mongo.Client
}
type Queue interface {
	GetAllJobs() ([]*model.Job, error)
	AddJob(*model.JobInput) (*string, error)
	ClaimJob() (*model.Job, error)
	FinishJob(jobID string) (*model.Job, error)
	GetJobsByStatus(status model.JobStatus) ([]*model.Job, error)
	GetJob(string) (*model.Job, error)
}
