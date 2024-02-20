package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"kwaaka-task/internal/domain"
	"kwaaka-task/pkg"
	"net/http"
	"time"
)

func (s *service) GetWeatherByCity(city string) (domain.Weather, error) {
	return s.repo.GetWeatherByCity(city)
}

func (s *service) UpdateWeather(city string) error {
	response, err := s.GetDataFromApi(city)
	if err != nil {
		return err
	}
	weather := domain.Weather{
		City:        response.City,
		Temperature: pkg.KelvinToCelsius(response.Main.Kelvin),
	}
	existingWeather, err := s.repo.GetWeatherByCity(city)
	if err != nil {
		if errors.Is(err, domain.ErrNoDocument) || existingWeather.City != weather.City {
			// fmt.Println("OK")

			weather.CreatedAt = time.Now()
			weather.UpdatedAt = time.Now()
			return s.repo.AddWeather(weather)
		}
		return errors.New(err.Error())
	}
	weather.UpdatedAt = time.Now()
	return s.repo.UpdateWeather(weather)
}

func (s *service) GetAllWeatherList() ([]domain.Weather, error) {
	return s.repo.GetAllWeatherList()
}

func (s *service) GetDataFromApi(city string) (domain.ApiResponse, error) {
	// pkg.InfoLog.Println("getResponseBody")
	fmt.Println(s.conf.Api.Url)
	body, err := getResponseBody(s.conf.Api.Url, city, s.conf.Api.Key)
	if err != nil {
		return domain.ApiResponse{}, err
	}

	var response domain.ApiResponse

	if err := json.Unmarshal(body, &response); err != nil {
		return domain.ApiResponse{}, fmt.Errorf("problems with unmarshalling response: %v", err)
	}
	return response, nil
}

func getResponseBody(url, name, key string) ([]byte, error) {
	fullUrl := fmt.Sprintf("%s?q=%s&appid=%s", url, name, key)
	fmt.Println(fullUrl)
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	client := new(http.Client)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to requesting: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}
	return body, nil
}
