package main

import (
	"github.com/Aliev20/Gin/internal/user"
	"github.com/Aliev20/Gin/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("name")
	w.Write([]byte(name))
}

func main() {
	logger := logging.GetLogger()

	logger.Println("Create router")
	router := httprouter.New()

	handler := user.New(logger)
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	logger := logging.GetLogger()

	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Infoln("Server started to: http://localhost:8000")
	logger.Fatal(server.Serve(listener))
}
