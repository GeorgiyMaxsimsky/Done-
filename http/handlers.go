package http

import (
	"done/todo"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type HTTPTaskHandlers struct {
	toDoList *todo.List
}

func NewHTTPhandlers(toDoList *todo.List) *HTTPTaskHandlers {
	return &HTTPTaskHandlers{
		toDoList: toDoList,
	}
}

/*
pattern: /tasks
method: POST
info: JSON  in HTTP request body

succeed:
-status code : 201
-responce body: JSON represent created task

failed:
- status code: 400,409,500
-responce body: JSON with error + time
*/
func (h *HTTPTaskHandlers) HandleCreateTask(w http.ResponseWriter, r *http.Request) {

	var taskDto TaskDto
	if err := json.NewDecoder(r.Body).Decode(&taskDto); err != nil {
		errDTO := ErrorDto{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDTO.toString(), http.StatusBadRequest)
		return

	}
	if err := taskDto.ValidateToCreate(); err != nil {
		errDTO := ErrorDto{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDTO.toString(), http.StatusBadRequest)
		return
	}

	//change when id generator be created

	toDoTask := todo.NewTask(1, taskDto.Title, taskDto.Description)
	if err := h.toDoList.AddTask(toDoTask); err != nil {
		errDTO := ErrorDto{
			Message: err.Error(),
			Time:    time.Now(),
		}
		if errors.Is(err, todo.ErrTaskAlreadyExist) {
			http.Error(w, errDTO.toString(), http.StatusConflict)
		} else {
			http.Error(w, errDTO.toString(), http.StatusInternalServerError)
		}
		return

	}

	b, err := json.MarshalIndent(toDoTask, "", "   ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http responce ", err)
		return
	}

}

/*
pattern: /tasks/{id}
method: GET
info: pattern

succeed:
-status code : 200 OK
-responce body: JSON represented found task

failed:
- status code: 400,404,409,500
-responce body: JSON with error + time
*/
func (h *HTTPTaskHandlers) HandleGetTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	task, err := h.toDoList.GetTask(id)
	if err != nil {
		errDto := NewErrDTO("Task is not found")
		if errors.Is(err, todo.ErrTaskNotFound) {

			http.Error(w, errDto.toString(), http.StatusNotFound)
		} else {
			http.Error(w, errDto.toString(), http.StatusInternalServerError)
		}
		return
	}
	b, err := json.MarshalIndent(task, "", "   ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)

}

/*
pattern: /tasks
method: GET
info:

succeed:
-status code : 200
-responce body: JSON represent tasks

failed:
- status code: 400,409,500
-responce body: JSON with error + time
*/
func (h *HTTPTaskHandlers) HandleGetAllTask(w http.ResponseWriter, r *http.Request) {
	tasks := h.toDoList.ListTasks()

	b, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {

		fmt.Println("failed to write http responce ", err)
		return
	}

}

/*
pattern: /tasks?complited=false
method: GET
info: querry params

succeed:
-status code : 200
-responce body: JSON represent tasks

failed:
- status code: 400,409,500
-responce body: JSON with error + time
*/
func (h *HTTPTaskHandlers) HandleGetAllUncomplitedTask(w http.ResponseWriter, r *http.Request) {
	tasks := h.toDoList.ListUnCompletedTasks()

	b, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {

		fmt.Println("failed to write http responce ", err)
		return
	}

}

/*
pattern: /tasks?complited=true
method: GET
info: querry params

succeed:
-status code : 200
-responce body: JSON represent tasks

failed:
- status code: 400,409,500
-responce body: JSON with error + time
*/

func (h *HTTPTaskHandlers) HandleGetAllComplitedTask(w http.ResponseWriter, r *http.Request) {
	tasks := h.toDoList.ListCompletedTasks()

	b, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {

		fmt.Println("failed to write http responce ", err)
		return
	}

}

/*
pattern: /tasks/{id}
method: PATCH
info: pattern + JSON in request body


succeed:
-status code : 200
-responce body: JSON represent changed task

failed:
- status code: 400,409,500
-responce body: JSON with error + time
*/

func (h *HTTPTaskHandlers) HandleCompleteTask(w http.ResponseWriter, r *http.Request) {
	var completedDto CompletedDto
	if err := json.NewDecoder(r.Body).Decode(&completedDto); err != nil {
		errDto := NewErrDTO("Inncorect format of completed task")

		http.Error(w, errDto.toString(), http.StatusBadRequest)
		return

	}

	id := mux.Vars(r)["id"]

	if completedDto.Completed {
		errDto := NewErrDTO("Task not found")
		if err := h.toDoList.CompleteTask(id); err != nil {
			if errors.Is(err, todo.ErrTaskNotFound) {
				http.Error(w, errDto.toString(), http.StatusNotFound)
			} else {
				http.Error(w, errDto.toString(), http.StatusNotFound)
			}
		}

	}

}

/*
pattern: /tasks/{id}
method: DELETE
info: pattern

succeed:
-status code : 204 No content
-responce body:

failed:
- status code: 400,404,409,500
-responce body: JSON with error + time
*/

func (h *HTTPTaskHandlers) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
}
