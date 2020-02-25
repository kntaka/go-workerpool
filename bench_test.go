package main

import (
	"testing"

	"github.com/kntaka/go-workerpool/job"
	"github.com/kntaka/go-workerpool/pool"
)

func BenchmarkConcurrent(b *testing.B) {
	collector := pool.StartDispatcher(workerCount) // start up worker pool

	for n := 0; n < b.N; n++ {
		for i, j := range job.CreateJobs(20) {
			collector.Work <- pool.Work{Job: j, ID: i}
		}
	}
}

func BenchmarkNonConcurrent(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, j := range job.CreateJobs(20) {
			job.DoWork(j, 1)
		}
	}
}
