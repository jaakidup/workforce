package workforce

// Job describes the task at hand
type Job struct {
	ID          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	WorkerID    string `json:"worker_id,omitempty"`
	Completed   bool   `json:"completed,omitempty"`
}

// JobList ...
type JobList struct {
	ID   string
	Jobs []Job
}
