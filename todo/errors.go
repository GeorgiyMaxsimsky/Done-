package todo

import "errors"

var ErrTaskNotFound = errors.New("Task are not found")

var ErrTaskAlreadyExist = errors.New("Task  with this id is already exist")
