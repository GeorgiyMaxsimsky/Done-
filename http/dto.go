package http

import "time"

type TaskDto struct {
	Title string `json:"title"`

	Description string `json:"description"`
}

type ErrorDto struct {
	Message string
	Time time.Time
}


func (e ErrorDto) toString 