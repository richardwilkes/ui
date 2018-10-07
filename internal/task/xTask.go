package task

import (
	"sync"
)

var (
	nextID uint64 = 1
	lock   sync.Mutex
	tasks  = make(map[uint64]func())
)

// Record the task to be executed. Returns the ID associated with it.
func Record(task func()) uint64 {
	if task == nil {
		panic("nil task not permitted")
	}
	lock.Lock()
	id := nextID
	nextID++
	tasks[id] = task
	lock.Unlock()
	return id
}

// Dispatch a task by ID.
func Dispatch(id uint64) {
	lock.Lock()
	task := tasks[id]
	if task != nil {
		delete(tasks, id)
	}
	lock.Unlock()
	if task != nil {
		task()
	}
}
