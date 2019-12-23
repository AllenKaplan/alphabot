package weather

import (
	"context"
	"github.com/AllenKaplan/alphabot/weather/handler"
	"github.com/briandowns/openweathermap"
)

type WeatherService struct {
	context.Context
}

func NewWeatherClient(ctx context.Context) WeatherService {
	return WeatherService{
		Context: ctx,
	}
}

type Weather interface {
	GetWeatherByLocation(string) (string, error)
}

func (s WeatherService) GetWeatherByLocation(req string) (*openweathermap.CurrentWeatherData, error) {
	res, err := weather.GetWeatherByLocation(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
