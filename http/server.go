package http

import "github.com/gorilla/mux"

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
}
