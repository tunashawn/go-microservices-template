package controllers

import (
	"go-microservices-template/internal/response"
	"go-microservices-template/pkg/service_A/services"
	"net/http"
)

type Controller interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

type ControllerImpl struct {
	response response.CustomResponse
	service  services.Service
}

func NewController() (Controller, error) {
	//service, err := services.NewService()
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &ControllerImpl{
	//	response: response.CustomResponse{},
	//	service:  service,
	//}, nil

	// TODO: uncomment above and delete below to use services and deeper level of code
	return &ControllerImpl{
		response: response.CustomResponse{},
		service:  nil,
	}, nil
}

func (c *ControllerImpl) Ping(w http.ResponseWriter, r *http.Request) {
	c.response.Success("pong", w)
}
