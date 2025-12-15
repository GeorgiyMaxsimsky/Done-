package todo

import (
	"time"
)

type Task struct {
	Id          int
	Title       string
	Description string

	CreatedTime time.Time
	ti
}
