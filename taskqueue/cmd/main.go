// Package cmd değil de main dememizin sebebi: Only package main can be compiled
// into an executable.

package main

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/Salih04/CLI-Task-Queue-with-Priority-Rules/internal/rules"
	"github.com/Salih04/CLI-Task-Queue-with-Priority-Rules/internal/task"
)

func main() {
	// All tasks starts from Pending

	taskA := &task.Task{
		ID:       "task-a",
		Name:     "salih",
		Priority: 10,
		Status:   task.Pending,
		Action: func() error {
			fmt.Println(" Connecting to database... ")
			time.Sleep(2 * time.Second)
			return nil
		},
	}
	taskB := &task.Task{
		ID:        "task-b",
		Name:      "ayşe",
		Priority:  8,
		Status:    task.Pending,
		DependsOn: []string{"task-a"},
		Action: func() error {
			fmt.Println(" Connecting to database...")
			time.Sleep(1 * time.Second)
			return nil
		},
	}
	taskC := &task.Task{
		ID:       "task-c",
		Name:     "mehmet",
		Priority: 6,
		Status:   task.Pending,
		Action: func() error {
			fmt.Println(" Connecting to database... ")
			time.Sleep(1 * time.Second)
			return errors.New("Schema Invalid!")
		},
	}
	taskD := &task.Task{
		ID:           "task-d",
		Name:         "meltem",
		Priority:     5,
		Status:       task.Pending,
		SkipIfFailed: []string{"task-c"},
		Action: func() error {
			fmt.Println(" seeding data... ")
			return nil
		},
	}

	allTasks := []*task.Task{taskA, taskB, taskC, taskD} // SLICE

	// WE can build this in 1 loop but for READABILITY and SEPERATION we build it seperately

	// Build registry - map each task ID to its task pointer
	registry := map[string]*task.Task{}
	for _, t := range allTasks {
		registry[t.ID] = t
	}

	for pendingCount(allTasks) > 0 {
		ready := []*task.Task{}

		for _, t := range allTasks {
			if t.Status != task.Pending {
				continue
			}
			result := rules.Evaluate(t, registry)
			if result.Skip {
				fmt.Printf("[SKIP] %s\n", t.Name)
				t.Status = task.Skipped
			} else if result.CanRun {
				ready = append(ready, t)
			} else {
				t.Status = task.Pending
			}
		}

		var wg sync.WaitGroup
		for _, t := range ready {
			wg.Add(1)
			go func(t *task.Task) {
				defer wg.Done()
				fmt.Printf("[RUN] %s\n", t.Name)
				err := t.Action()
				if err != nil {
					fmt.Printf("[FAIL] %s - %v\n", t.Name, err)
					t.Status = task.Failed
				} else {
					t.Status = task.Completed
				}
			}(t)
		}
		wg.Wait()
	}
	fmt.Println("\n===== RESULTS =====")
	for _, t := range allTasks {
		fmt.Printf(" %s -> %v\n", t.Name, t.Status)
	}

	/*

		round 1:
			loop through allTasks
				evaluate each one
				if CanRun  → add to ready slice
				if Skip    → mark as Skipped
				if Wait    → leave it, we'll check next round

			launch all ready tasks as goroutines simultaneously
			wait for all of them to finish

		round 2:
			same thing — now taskB sees taskA is Completed, so it becomes ready
			...repeat until nothing is Pending

	*/
}

func pendingCount(tasks []*task.Task) int {
	count := 0
	for _, t := range tasks {
		if t.Status == task.Pending {
			count++
		}
	}
	return count
}
