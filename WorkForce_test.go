package workforce

import (
	"fmt"
	"testing"
)

func TestWorkforce(t *testing.T) {
	joblist := []Job{
		Job{
			ID:          "1",
			Description: "Test Job",
		},
		Job{
			ID:          "2",
			Description: "Test Job",
		},
	}

	workforce := Workforce{}
	task := func(worker *Worker) {
		worker.Job.WorkerID = worker.ID
		// do some random stuff
		worker.Job.Completed = true
	}

	completed := workforce.StartSweatshop(2, joblist, task)

	fmt.Println(completed.Jobs)
	//Output: [{1 Test Job IDI-0 true} {2 Test Job IDI-1 true}]
	if !completed.Jobs[0].Completed == true {
		t.Error("Should have been completed")
	}
	if !completed.Jobs[1].Completed == true {
		t.Error("Should have been completed")
	}
}

func ExampleSample() {
	Sample()
	//Output: Sample output!

}

func Example() {

	joblist := []Job{
		Job{
			ID:          "1",
			Description: "Test Job",
		},
		Job{
			ID:          "2",
			Description: "Test Job",
		},
	}

	task := func(worker *Worker) {
		worker.Job.WorkerID = worker.ID
		// do some random stuff...
		worker.Job.Completed = true
	}
	workforce := &Workforce{}
	completed := workforce.StartSweatshop(2, joblist, task)

	fmt.Println(completed.Jobs)
	//Output: [{1 Test Job IDI-0 true} {2 Test Job IDI-1 true}]

}
