package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type JobStatus string

const (
	Waiting   JobStatus = "Waiting"
	Resources JobStatus = "Not Enough Resources"
	Fulfilled JobStatus = "Fulfilled"
)

type job struct {
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
	GetAll() []job
}

func (q QueueModel) GetAll() []job {
	queueCollection := q.DB.Database("Jobs").Collection("Queue")
	results, err := queueCollection.Find(context.TODO(), bson.D{})

	if err != nil {
		fmt.Errorf("Error retrieving jobs: %w", err)
	}

	var jobResult *[]job = new([]job)
	if err = results.All(context.TODO(), jobResult); err != nil {
		panic(err)
	}

	return *jobResult
}

func (q QueueModel) AddNewJob(jobToInsert *job) {
	queueCollection := q.DB.Database("Jobs").Collection("Queue")
	_, err := queueCollection.InsertOne(context.TODO(), jobToInsert)

	if err != nil {
		panic(err)
	}
}
