package todo

import "time"

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todo []item

func (t *Todo) Add(task string) {
	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},		
	}

	*t = append(*t, todo)
}