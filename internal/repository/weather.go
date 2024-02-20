package repository

import (
	"context"
	"errors"
	"kwaaka-task/internal/domain"
	"kwaaka-task/pkg"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *repositoryMongo) AddWeather(weather domain.Weather) error {
	_, err := r.collection.InsertOne(context.Background(), weather)
	if err != nil {
		pkg.ErrorLog.Printf("Error in creating: %v", err)
		return err
	}
	return nil
}

func (r *repositoryMongo) GetWeatherByCity(city string) (domain.Weather, error) {
	var weather domain.Weather

	filter := bson.D{{"city", city}}

	if err := r.collection.FindOne(context.Background(), filter).Decode(&weather); err != nil {
		pkg.ErrorLog.Printf("Error in getting session document: %v", err)
		return weather, errors.New("Error no documents")
	}

	return weather, nil
}

func (r *repositoryMongo) UpdateWeather(weather domain.Weather) error {
	filter := bson.D{{"city", weather.City}}
	update := bson.D{{"$set", bson.D{
		{"temperature", weather.Temperature},
		{"updated_at", weather.UpdatedAt},
	}}}

	if _, err := r.collection.UpdateOne(context.Background(), filter, update); err != nil {
		pkg.ErrorLog.Printf("Error in updating session document: %v", err)
		return errors.New("Can`t update session document")
	}
	return nil
}
