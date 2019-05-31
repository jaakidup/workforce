package workforce

import (
	"fmt"
	"strconv"
)

// Workforce was created to simplify the createion of a multi go routine section in a system
type Workforce struct {
	Delegator
	Worker
	Receiver
	Task func(*Worker)
}

// StartSweatshop was created to ease the creation of a multi goroutine system
//
// @param numberOfWorkers int that gets activated for the job
//
// @param jobList []Job all the Jobs that need to be done.
func (workforce *Workforce) StartSweatshop(numberOfWorkers int, jobList []Job, task func(*Worker)) JobList {

	// Create all the comunication channels
	// Start the receiver, listen for completed jobs
	// Start the workers, listen for Job
	// Start the delegator, send Jobs to workers

	delegatorToWorker, workerToReceiver := make(JobChan), make(JobChan, numberOfWorkers)
	jobsListChan := make(JobsListChan)

	// Receives Jobs from individual Workers
	// go receiver(workerToReceiver)
	receiver := NewReceiver(workerToReceiver, jobsListChan, len(jobList))
	go receiver.Receive()

	for i := 0; i < numberOfWorkers; i++ {
		workerID := "IDI-" + strconv.Itoa(i)
		worker := NewWorker(workerID, delegatorToWorker, workerToReceiver)
		worker.Task = task
		go worker.Ready()
	}

	// Sends the jobs to the Workers
	go NewDelegator(delegatorToWorker, jobList).Delegate()

	// Now we just receive the completed joblists
	completed := <-jobsListChan
	return completed
}

// Sample is purely for testing the examples of godoc
func Sample() {
	fmt.Println("Sample output!")
}
