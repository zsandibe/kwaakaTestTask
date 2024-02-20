package service

import (
	"encoding/json"
	"fmt"
	"io"
	"kwaaka-task/internal/domain"
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
		Temperature: response.Main.Kelvin,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	existingWeather, err := s.repo.GetWeatherByCity(city)
	if err != nil {
		return err
	}
	if existingWeather.City != weather.City {
		return s.repo.AddWeather(weather)
	}

	return s.repo.UpdateWeather(weather)
}

func (s *service) GetDataFromApi(city string) (domain.ApiResponse, error) {

	// pkg.InfoLog.Println("getResponseBody")
	body, err := getResponseBody(s.conf.Api.Url, city)
	if err != nil {
		return domain.ApiResponse{}, err
	}
	var response domain.ApiResponse

	if err := json.Unmarshal(body, &response); err != nil {
		return domain.ApiResponse{}, fmt.Errorf("problems with unmarshalling response: %v", err)
	}

	return response, nil
}

func getResponseBody(url, name string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	query := req.URL.Query()
	query.Add("q", name)

	req.URL.RawQuery = query.Encode()

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
