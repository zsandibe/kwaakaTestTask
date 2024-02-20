package service

import (
	"kwaaka-task/config"
	"kwaaka-task/internal/domain"
	"kwaaka-task/internal/repository"
)

type Service interface {
	GetWeatherByCity(city string) (domain.Weather, error)
	UpdateWeather(city string) error
	GetDataFromApi(city string) (domain.ApiResponse, error)
}

type service struct {
	conf config.Config
	repo repository.Repository
}

func NewService(repo repository.Repository, conf config.Config) *service {
	return &service{
		repo: repo,
		conf: conf,
	}
}
