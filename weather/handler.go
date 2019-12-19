package weather

type Service struct{
}

func NewWeatherService() Service {
	return Service{}
}

type Weather interface {
	GetWeatherByLocation(string) (string, error)
}

func (s Service) GetWeatherByLocation(req string) (string, error) {
	res, err := GetWeatherByLocation(req)
	if err != nil {
		return "", err
	}

	return res, nil
}
