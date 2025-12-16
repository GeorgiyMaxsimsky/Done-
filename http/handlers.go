package http

import (
	"done/todo"
	"encoding/json"
	"net/http"
	"time"
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
		http.Error(w)

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
