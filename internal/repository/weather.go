package repository

import (
	"context"
	"errors"
	"fmt"
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
	fmt.Println("Updating weather")

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

func (r *repositoryMongo) GetAllWeatherList() ([]domain.Weather, error) {
	ctx := context.Background()
	rows, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.New("Can`t get all sessions rows")
	}
	defer rows.Close(ctx)

	var weatherList []domain.Weather

	for rows.Next(ctx) {
		var weather domain.Weather
		if err := rows.Decode(&weather); err != nil {
			fmt.Println("Error in getting")
			return nil, errors.New("Failed to decode row")
		}
		weatherList = append(weatherList, weather)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return weatherList, nil
}
