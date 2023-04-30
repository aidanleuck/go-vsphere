package queue

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (q *QueueModel) GetAllJobs() (*[]Job, error) {
	queueCollection := q.DB.Database("Jobs").Collection("Queue")

	// If the queue is empty don't worry about it.
	results, err := queueCollection.Find(context.TODO(), bson.D{})

	if err != nil {
		fmt.Errorf("Error retrieving jobs: %w", err)
		return nil, err
	}

	var jobResult *[]Job = new([]Job)
	if err = results.All(context.Background(), jobResult); err != nil {
		// If we can't even return an empty collection something is wrong. Panic.
		panic(err)
	}

	return jobResult, nil
}

func (q *QueueModel) GetJobsByStatus(status JobStatus) (*[]Job, error) {
	queueCollection := q.DB.Database("Jobs").Collection("Queue")
	results, err := queueCollection.Find(context.Background(), bson.D{{"Status", status}})

	if err != nil {
		fmt.Errorf("Error retrieving jobs: %w", err)
		return nil, err
	}

	var jobResult *[]Job = new([]Job)
	if err = results.All(context.Background(), jobResult); err != nil {
		fmt.Errorf("Error decoding job: %w", err)
		return nil, err
	}

	return jobResult, nil
}

func (q *QueueModel) GetJobByStatus(status JobStatus) (*Job, error) {
	queueCollection := q.DB.Database("Jobs").Collection("Queue")
	var result Job
	err := queueCollection.FindOne(context.Background(), bson.D{{"Status", status}}).Decode(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (q *QueueModel) ClaimJob() (*Job, error) {
	queueCollection := q.DB.Database("Jobs").Collection("Queue")
	var result Job

	// Sorts descending by priority
	sortOptions := options.FindOneAndUpdate().SetSort(bson.D{{"Priority", -1}})
	update := bson.M{
		"$set": bson.M{"Status": Fulfilled},
	}
	filter := bson.D{{"Status", Waiting}}
	err := queueCollection.FindOneAndUpdate(context.Background(), filter, update, sortOptions).Decode(&result)

	if err != nil {
		return nil, nil
	}

	return &result, nil
}

func (q *QueueModel) FinishJob(jobID *Job) error {
	queueCollection := q.DB.Database("Jobs").Collection("Queue")
	mongoFilter := bson.D{{"JobID", jobID}}
	result := queueCollection.FindOneAndDelete(context.Background(), mongoFilter)

	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

func (q QueueModel) AddJob(jobToInsert *Job) error {
	queueCollection := q.DB.Database("Jobs").Collection("Queue")
	_, err := queueCollection.InsertOne(context.Background(), jobToInsert)

	if err != nil {
		return err
	}
	return nil
}

func (q QueueModel) GetJob(jobID string) (*Job, error) {
	queueCollection := q.DB.Database("Jobs").Collection("Queue")
	mongoFilter := bson.D{{"JobID", jobID}}
	foundJob := queueCollection.FindOne(context.Background(), mongoFilter)
	var decodedJob Job
	if foundJob.Err() != nil {
		return nil, foundJob.Err()
	}

	foundJob.Decode(&decodedJob)
	return &decodedJob, nil
}
