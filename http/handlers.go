package http

import (
	"done/todo"
	"net/http"
)

type HTTPHandlers struct {
	toDoList *todo.List
}

func NewHTTPhandlers(toDoList *todo.List) *HTTPHandlers {
	return &HTTPHandlers{
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
func (h *HTTPHandlers) HandleCreateTask(w http.ResponseWriter, r *http.Request) {

}

/*
pattern: /tasks/{id}
method: GET
info: pattern

succeed:
-status code : 200 OK
-responce body: JSON represent found task

failed:
- status code: 400,404,409,500
-responce body: JSON with error + time
*/
func (h *HTTPHandlers) HandleGetTask(w http.ResponseWriter, r *http.Request) {

}
