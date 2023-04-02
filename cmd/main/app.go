package main

import (
	"fmt"
	"github.com/Aliev20/Gin/internal/config"
	"github.com/Aliev20/Gin/internal/user"
	"github.com/Aliev20/Gin/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"time"
)

func main() {
	logger := logging.GetLogger()

	logger.Println("Create router")
	router := httprouter.New()

	cfg := config.GetConfig()

	handler := user.New(logger)
	handler.Register(router)

	start(router, cfg)
}

func start(router *httprouter.Router, config *config.Config) {
	logger := logging.GetLogger()

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", config.Listen.BindIp, config.Listen.Port))
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Infoln("Server started to: http://" + fmt.Sprintf("%s:%s", config.Listen.BindIp, config.Listen.Port))
	logger.Fatal(server.Serve(listener))
}
