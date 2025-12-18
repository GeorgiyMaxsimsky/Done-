package http

import (
	"encoding/json"
	"errors"

	"time"
)

type TaskDto struct {
	Title string `json:"title"`

	Description string `json:"description"`
}

func (t TaskDto) ValidateToCreate() error {
	if t.Title == "" {
		return errors.New("title is empty")

	}
	if t.Description == "" {
		return errors.New("description is empry")
	}
	return nil

}

type CompletedDto struct {
	Completed bool
}

func NewCompletedDto(isit bool) CompletedDto {
	return CompletedDto{
		Completed: isit,
	}
}

type ErrorDto struct {
	Message string
	Time    time.Time
}

func NewErrDTO(message string) ErrorDto {

	return ErrorDto{
		Message: message,
		Time:    time.Now(),
	}

}

func (e ErrorDto) toString() string {
	b, err := json.MarshalIndent(e, "", "	")

	if err != nil {
		panic(err)
	}
	return string(b)
}
