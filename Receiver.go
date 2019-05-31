package workforce

// NewReceiver  receveives completed jobs
func NewReceiver(receiveChan JobChan, jobsOutChan JobsListChan, totalJobs int) *Receiver {
	return &Receiver{ReceiveChan: receiveChan, SendJobsListChan: jobsOutChan, TotalJobs: totalJobs}
}

// Receiver recieves Job from Workers
type Receiver struct {
	ReceiveChan      JobChan
	SendJobsListChan JobsListChan
	Jobs             []Job
	TotalJobs        int
}

// Receive set the receiver to listening for completed Jobs
func (receiver *Receiver) Receive() {
	for {
		receiver.Jobs = append(receiver.Jobs, <-receiver.ReceiveChan)
		if len(receiver.Jobs) == receiver.TotalJobs {
			break
		}
	}
	jobsList := JobList{
		ID:   "some-ID",
		Jobs: receiver.Jobs,
	}
	receiver.SendJobsListChan <- jobsList
}
