package delivery

import "kwaaka-task/internal/service"

type handler struct {
	service service.Service
}

func NewHandler(service service.Service) *handler {
	return &handler{service: service}
}
