package queue

import (
	"context"
	"fmt"
	"time"

	"vsphere_module/src/graph/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (q *QueueModel) GetAllJobs() ([]*model.Job, error) {
	queueCollection := q.DB.Database("Jobs").Collection("Queue")

	// If the queue is empty don't worry about it.
	results, err := queueCollection.Find(context.TODO(), bson.D{})

	if err != nil {
		fmt.Errorf("Error retrieving jobs: %w", err)
		return nil, err
	}

	var jobResult []*model.Job
	if err = results.All(context.Background(), &jobResult); err != nil {
		// If we can't even return an empty collection something is wrong. Panic.
		panic(err)
	}

	return jobResult, nil
}

func (q *QueueModel) GetJobsByStatus(status model.JobStatus) ([]*model.Job, error) {
	queueCollection := q.DB.Database("Jobs").Collection("Queue")
	results, err := queueCollection.Find(context.Background(), bson.D{{"status", status}})

	if err != nil {
		fmt.Errorf("Error retrieving jobs: %w", err)
		return nil, err
	}

	var jobResult []*model.Job
	if err = results.All(context.Background(), &jobResult); err != nil {
		fmt.Errorf("Error decoding job: %w", err)
		return nil, err
	}

	return jobResult, nil
}

func (q *QueueModel) ClaimJob() (*model.Job, error) {
	queueCollection := q.DB.Database("Jobs").Collection("Queue")
	var result model.Job

	// Sorts descending by priority
	sortOptions := options.FindOneAndUpdate().SetSort(bson.D{{"priority", -1}})
	update := bson.M{
		"$set": bson.M{"status": model.JobStatusFulfilled},
	}
	filter := bson.D{{"status", model.JobStatusCreating}}
	err := queueCollection.FindOneAndUpdate(context.Background(), filter, update, sortOptions).Decode(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (q *QueueModel) FinishJob(jobID string) (*model.Job, error) {
	queueCollection := q.DB.Database("Jobs").Collection("Queue")
	mongoFilter := bson.D{{"jobid", jobID}}
	result := queueCollection.FindOneAndDelete(context.Background(), mongoFilter)
	if result.Err() != nil {
		return nil, result.Err()
	}
	var finishedJob *model.Job
	err := result.Decode(&finishedJob)
	if err != nil {
		return nil, err
	}

	return finishedJob, nil
}

func generateJobID(input *model.JobInput) string {
	now := time.Now().Unix()
	return fmt.Sprintf("%s-%s-%d@%d", input.Product, input.BranchName, input.BuildNumber, now)
}

func (q QueueModel) AddJob(jobToInsert *model.JobInput) (*string, error) {
	newJobID := generateJobID(jobToInsert)
	var priority int
	if nil != jobToInsert.Priority {
		priority = *jobToInsert.Priority
	} else {
		priority = 0
	}
	job := &model.Job{
		JobID:       newJobID,
		BranchName:  jobToInsert.BranchName,
		BuildNumber: jobToInsert.BuildNumber,
		Product:     jobToInsert.Product,
		TotalVM:     jobToInsert.TotalVM,
		RemainingVM: jobToInsert.TotalVM,
		Status:      model.JobStatusNew,
		Priority:    priority,
	}
	queueCollection := q.DB.Database("Jobs").Collection("Queue")
	_, err := queueCollection.InsertOne(context.Background(), job)
	if err != nil {
		return nil, err
	}
	return &newJobID, nil
}

func (q QueueModel) GetJob(jobID string) (*model.Job, error) {
	queueCollection := q.DB.Database("Jobs").Collection("Queue")
	mongoFilter := bson.D{{"jobid", jobID}}
	foundJob := queueCollection.FindOne(context.Background(), mongoFilter)
	var decodedJob model.Job
	if foundJob.Err() != nil {
		return nil, foundJob.Err()
	}

	foundJob.Decode(&decodedJob)
	return &decodedJob, nil
}
