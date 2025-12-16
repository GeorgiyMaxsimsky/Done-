package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	httpHandlers *HTTPTaskHandlers
}

func NewServer(http *HTTPTaskHandlers) *HTTPServer {
	return &HTTPServer{
		httpHandlers: http,
	}
}

func (s *HTTPServer) startServer() error {
	router := mux.NewRouter()
	router.Path("/tasks").Methods("POST").HandlerFunc(s.httpHandlers.HandleCreateTask)
	router.Path("/task/{id}").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetTask)
	router.Path("/task").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetAllTask)
	router.Path("/task").Methods("GET").Queries("complited", "true").HandlerFunc(s.httpHandlers.HandleGetAllUncomplitedTask)
	router.Path("/task").Methods("GET").Queries("complited", "false").HandlerFunc(s.httpHandlers.HandleGetAllComplitedTask)
	router.Path("/task/{id}").Methods("PATCH").HandlerFunc(s.httpHandlers.HandleCompleteTask)
	router.Path("/task/{id}").Methods("DELETE").HandlerFunc(s.httpHandlers.HandleDeleteTask)
	return http.ListenAndServe(":9091", router)
}
