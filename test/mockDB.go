package test

import "vsphere_module/src/queue"

type mockQueueModel struct{}

func (q *mockQueueModel) GetAllJobs() []queue.Job {
	var jobs []queue.Job

	jobs = append(jobs, queue.Job{BranchName: "Hello"})
	return jobs
}

func (q *mockQueueModel) AddJob(jobToInsert *queue.Job) {

}
