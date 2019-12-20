package weather

import (
	"context"
	"github.com/AllenKaplan/alphabot/weather/handler"
)

type WeatherService struct{
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

func (s WeatherService) GetWeatherByLocation(req string) (string, error) {
	res, err := weather.GetWeatherByLocation(req)
	if err != nil {
		return "", err
	}

	return res, nil
}
