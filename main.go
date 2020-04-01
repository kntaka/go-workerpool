package main

import (
	"log"

	"github.com/knwoop/go-workerpool/job"
	"github.com/knwoop/go-workerpool/pool"
)

const (
	workerCount = 5
	jobCount    = 100
)

func main() {
	log.Println("starting application...")
	collector := pool.StartDispatcher(workerCount) // start up worker pool

	for i, j := range job.CreateJobs(jobCount) {
		collector.Work <- pool.Work{Job: j, ID: i}
	}
}
