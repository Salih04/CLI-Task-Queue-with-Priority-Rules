# CLI Task Queue with Priority & Rules

A beginner-friendly Go project for learning backend concepts through a small task runner.

This project simulates a task queue where tasks can have priorities, dependencies, skip rules, and concurrent execution using goroutines.

---

# What This Project Teaches

- Go project structure
- Structs and methods
- `iota` enums
- Interfaces
- Priority queue with `container/heap`
- Maps and slices
- Error handling
- Goroutines
- `sync.WaitGroup`
- Basic rule engine logic

---

# Project Structure

```txt
CLI-Task-Queue-with-Priority-Rules/
│
├── go.mod
│
├── cmd/
│   └── main.go
│
└── internal/
    ├── task/
    │   └── task.go
    │
    ├── queue/
    │   └── queue.go
    │
    └── rules/
        └── rules.go
```

---

# How It Works

Each task has:

- ID
- Name
- Priority
- Status
- DependsOn
- SkipIfFailed
- Action

Example logic:

- Task B waits until Task A is completed
- Task D is skipped if Task C fails
- Independent tasks run concurrently using goroutines

---

# Task Statuses

```go
Pending
Running
Completed
Failed
Skipped
```

These statuses are defined using Go's `iota`.

---

# Running The Project

```bash
go run cmd/main.go
```

---

# Example Output

```txt
[RUN] mehmet
 Connecting to database...

[RUN] salih
 Connecting to database...

[FAIL] mehmet - Schema Invalid!

[SKIP] meltem

[RUN] ayşe
 Connecting to database...
```

---

# Current Features

- Task definition system
- Priority queue structure
- Dependency rule evaluation
- Skip-if-failed logic
- Concurrent task execution
- Final task status summary

---

# Concepts Learned During Development

## Structs
Used to model tasks and task states.

## `iota`
Used for enum-like task statuses.

## Heap Interface
Implemented Go's `heap.Interface` to build a priority queue.

## Maps
Used as a registry for fast task lookup by ID.

## Goroutines
Used to execute independent tasks concurrently.

## WaitGroup
Used to wait until all concurrent tasks finish.

## Error Handling
Go returns errors as values instead of exceptions.

---

# Future Improvements

- Save task state to JSON
- Load task state from JSON
- Add CLI commands
- Add unit tests
- Add better logging
- Add circular dependency detection
- Add retry logic for failed tasks

---

# Why I Built This

I built this project while learning Go in a project-based way.

The goal was to understand backend engineering concepts like:

- queues
- task execution
- concurrency
- dependency management
- rule engines

through a simple but realistic project.