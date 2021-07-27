package handlers

import (
	"net/http"

	s "github.com/khihadysucahyo/go-echo-boilerplate/server"

	"github.com/khihadysucahyo/go-echo-boilerplate/responses"

	"github.com/khihadysucahyo/go-echo-boilerplate/requests"

	"github.com/khihadysucahyo/go-echo-boilerplate/repositories"

	"github.com/khihadysucahyo/go-echo-boilerplate/models"

	"github.com/labstack/echo/v4"
)

type RegisterHandler struct {
	server *s.Server
}

func NewRegisterHandler(server *s.Server) *RegisterHandler {
	return &RegisterHandler{server: server}
}

func (registerHandler *RegisterHandler) Register(c echo.Context) error {
	registerRequest := new(requests.RegisterRequest)

	if err := c.Bind(registerRequest); err != nil {
		return err
	}

	if err := registerRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid")
	}

	existUser := models.User{}
	userRepository := repositories.NewUserRepository(registerHandler.server.DB)
	userRepository.GetUserByEmail(&existUser, registerRequest.Email)

	if existUser.ID != 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "User already exists")
	}

	if err := userRepository.Register(registerRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Server error")
	}

	return responses.MessageResponse(c, http.StatusCreated, "User successfully created")
}
