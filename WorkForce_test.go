package workforce

import (
	"fmt"
	"testing"
)

var joblist = []Job{
	Job{
		ID: "2",
		Input: map[string]interface{}{
			"method": "GET",
			"url":    "http://someurl.com",
			"body": map[string]string{
				"testdata1": "sample",
				"testdata2": "sample",
			},
		},
	},
}

func TestWorkforce(t *testing.T) {

	workforce := Workforce{}
	task := func(worker *Worker) {
		// for name, value := range worker.Job.Input {
		// 	fmt.Println(name, " => ", value)
		// }
		worker.Job.Output = "Sample Data"

		worker.Job.Completed = true
	}

	completed := workforce.StartSweatshop(2, joblist, task)

	data := completed.Jobs[0].Output
	if data != "Sample Data" {
		t.Error("Should have been 'Sample Data' ")
	}
	if !completed.Jobs[0].Completed == true {
		t.Error("Should have been completed")
	}
}

func ExampleSample() {
	Sample()
	//Output: Sample output!

}

func Example() {

	task := func(worker *Worker) {
		// for name, value := range worker.Job.Input {
		// 	fmt.Println(name, " => ", value)
		// }
		worker.Job.Output = "Sample Data"
		worker.Job.Completed = true
	}
	workforce := &Workforce{}
	completed := workforce.StartSweatshop(2, joblist, task)

	fmt.Println(completed.Jobs)
	//Output: [{2 map[body:map[testdata1:sample testdata2:sample] method:GET url:http://someurl.com] Sample Data  true}]

}
