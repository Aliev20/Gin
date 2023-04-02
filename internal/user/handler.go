package user

import (
	"github.com/Aliev20/Gin/internal/handlers"
	"github.com/Aliev20/Gin/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct {
	logger *logging.Logger
}

func New(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetList)
	router.POST(usersURL, h.CreateUser)
	router.GET(userURL, h.GetUserByUUID)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartiallyUpdateUser)
	router.DELETE(userURL, h.DeleteUser)
}

func (h *handler) GetList(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Write([]byte("GetList"))
}

func (h *handler) CreateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Write([]byte("CreateUser"))

}
func (h *handler) GetUserByUUID(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Write([]byte("GetUserByUUID"))

}
func (h *handler) UpdateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Write([]byte("UpdateUser"))

}
func (h *handler) PartiallyUpdateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Write([]byte("PartiallyUpdateUser"))

}
func (h *handler) DeleteUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Write([]byte("DeleteUser"))

}
