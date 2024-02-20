package repository

import (
	"kwaaka-task/config"
	"kwaaka-task/internal/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GetWeatherByCity(city string) (domain.Weather, error)
	UpdateWeather(weather domain.Weather) error
	AddWeather(weather domain.Weather) error
}

type repositoryMongo struct {
	collection *mongo.Collection
	conf       config.Config
}

func NewRepository(db *mongo.Database, conf config.Config) *repositoryMongo {
	return &repositoryMongo{
		collection: db.Collection(conf.Database.NameCollection),
	}
}
