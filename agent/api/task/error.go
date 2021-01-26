package task

// taskError indicates any error in the task
type Error struct {
	ErrorField   string `json:"ErrorField,omitempty"`
	ErrorCode    string `json:"ErrorCode,omitempty"`
	ErrorMessage string `json:"ErrorMessage,omitempty"`
	StatusCode   int    `json:"StatusCode,omitempty"`
	ResourceARN  string `json:"ResourceARN,omitempty"`
}

func (err Error) ErrorName() string {
	return "TaskError"
}

func (err Error) Error() string {
	return err.ErrorMessage
}
