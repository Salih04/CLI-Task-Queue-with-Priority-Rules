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
