package todo

import (
	"time"
)

type Task struct {
	Id          int
	Title       string
	Description string
	IsCompleted bool

	CreatedTime time.Time
	DoneAt      *time.Time
}

func NewTask(
	id int,
	title string,
	description string,

) Task {
	return Task{
		Id:          id,
		Title:       title,
		Description: description,
		IsCompleted: false,

		CreatedTime: time.Now(),
		DoneAt:      nil,
	}
}

func (task *Task) Complete() {

	doneTime := time.Now()

	task.IsCompleted = true
	task.DoneAt = &doneTime
}
