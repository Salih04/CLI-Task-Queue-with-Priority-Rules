package rules

import "github.com/Salih04/CLI-Task-Queue-with-Priority-Rules/internal/task"

type EvalResult struct {
	CanRun bool
	Skip   bool
	Reason string
}

func Evaluate(t *task.Task, registry map[string]*task.Task) EvalResult {
	for _, id := range t.DependsOn {
		dep := registry[id]
		if dep.Status != task.Completed {
			return EvalResult{CanRun: false}
		}
	}
	for _, id := range t.SkipIfFailed {
		dep := registry[id]
		if dep.Status == task.Failed {
			return EvalResult{Skip: true}
		}
		// If dep hasn't finished yet, wait - don't skip prematurely
		if dep.Status == task.Pending || dep.Status == task.Running {
			return EvalResult{CanRun: false}
		}
	}
	return EvalResult{CanRun: true} 

}
