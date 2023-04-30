package queue

import "go.mongodb.org/mongo-driver/mongo"

type JobStatus string

const (
	Waiting   JobStatus = "Waiting"
	Resources JobStatus = "Not Enough Resources"
	Fulfilled JobStatus = "Fulfilled"
)

type Job struct {
	JobID       string    `bson:"JobID"`
	BranchName  string    `bson:"BranchName"`
	BuildNumber uint64    `bson:"BuildNumber"`
	Product     string    `bson:"Product"`
	TotalVM     uint64    `bson:"TotalVM"`
	RemainingVM uint64    `bson:"RemainingVM"`
	Status      JobStatus `bson:"Status"`
	Priority    int       `bson:"Priority"`
}

type QueueModel struct {
	DB *mongo.Client
}
type Queue interface {
	GetAllJobs() *[]Job
	AddJob(*Job)
	ClaimJob() *Job
	FinishJob(*Job)
	GetJobsByStatus(*Job)
	GetJob(string) (*Job, error)
}
