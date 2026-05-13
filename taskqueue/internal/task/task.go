package task

type Status int

const (
	Pending Status = iota
	Running
	Completed
	Failed
	Skipped
)

type Task struct {
	ID           string
	Name         string
	Priority     int
	Status       Status
	DependsOn    []string
	SkipIfFailed []string
	Action       func() error
}

type TaskState struct {
	ID           string   `json:"id"` // these are struct tags they tell GO's JSON library what names to use in the JSON file.
	Name         string   `json:"name"`
	Priority     string   `json:"priority"`
	Status       Status   `json:"status"`
	DependsOn    []string `json:"depends_on"`
	SkipIfFailed []string `json:"skip_if_failed"`
}

func (s Status) String() string {
	switch s {
	case Pending:
		return "PENDING"
	case Completed:
		return "COMPLETED"
	case Failed:
		return "FAILED"
	case Skipped:
		return "SKIPPED"
	default:
		return "UNKOWN"
	}
}
