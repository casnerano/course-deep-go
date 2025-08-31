package goroutines_and_scheduler

import (
	"container/heap"
	"fmt"
)

type Tasks []Task

func (t Tasks) Len() int {
	return len(t)
}

func (t Tasks) Less(i, j int) bool {
	return t[i].Priority > t[j].Priority
}

func (t Tasks) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t *Tasks) Push(x any) {
	*t = append(*t, x.(Task))
}

func (t *Tasks) Pop() any {
	item := (*t)[len(*t)-1]
	*t = (*t)[:len(*t)-1]
	return item
}

type Task struct {
	Identifier int
	Priority   int
}

type Scheduler struct {
	tasks Tasks
}

func NewScheduler() Scheduler {
	tasks := make(Tasks, 0)
	heap.Init(&tasks)

	return Scheduler{
		tasks: tasks,
	}
}

func (s *Scheduler) AddTask(task Task) {
	heap.Push(&s.tasks, task)
}

func (s *Scheduler) ChangeTaskPriority(taskID int, newPriority int) {
	for index := range s.tasks {
		if s.tasks[index].Identifier == taskID {
			s.tasks[index].Priority = newPriority
			heap.Fix(&s.tasks, index)
			return
		}
	}
}

func (s *Scheduler) GetTask() Task {
	if len(s.tasks) == 0 {
		return Task{}
	}

	return heap.Pop(&s.tasks).(Task)
}

func (s *Scheduler) Print() {
	for _, t := range s.tasks {
		fmt.Print(t.Identifier, "(", t.Priority, ") ")
	}
	fmt.Println()
}
