package todo

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id          string
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
		Id:          uuid.NewString(),
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
