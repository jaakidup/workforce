package workforce

// NewWorker creates a new worker with send and receive channels
func NewWorker(ID string, receiveChan JobChan, sendChan JobChan) *Worker {
	return &Worker{
		ID:             ID,
		ReceiveJobChan: receiveChan,
		SendJobChan:    sendChan,
	}
}

// Worker was created to perform the Job as received from the Delegator
type Worker struct {
	ID             string
	ReceiveJobChan JobChan
	SendJobChan    JobChan
	Job            Job
	Task           func(*Worker)
}

// Ready performs the Job reciever from the delegator and sends it to receiver upon completion
func (worker *Worker) Ready() {
	worker.Job = <-worker.ReceiveJobChan
	worker.Task(worker)
	worker.SendJobChan <- worker.Job
}
