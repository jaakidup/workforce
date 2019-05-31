package workforce

// NewDelegator sets up the new delegator
// Returns *Delegator with jobs and send chan set up
func NewDelegator(send JobChan, jobList []Job) *Delegator {
	return &Delegator{sendChan: send, Jobs: jobList}
}

// Delegator was created to hand out the Jobs to Workers
type Delegator struct {
	sendChan JobChan
	Jobs     []Job
}

// Delegate send the Job to Worker
func (delegator *Delegator) Delegate() {
	for _, job := range delegator.Jobs {
		delegator.sendChan <- job
	}
}
