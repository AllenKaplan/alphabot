package weather

import (
	"fmt"
	"strings"

	owm "github.com/briandowns/openweathermap"
)

const (
	openWeatherApiKey = "b7d52c8f41df2cf2606593af4cf60778"
)

func GetWeatherByLocation(location string) (string, error) {

	const defaultFeeling = "I have no idea"
	feeling := defaultFeeling

	w, err := owm.NewCurrent("C", "EN", openWeatherApiKey) // (internal - OpenWeatherMap reference for kelvin) with English output
	if err != nil {
		return feeling, err
	}

	if err := w.CurrentByName(location); err != nil {
		return feeling, err
	}

	temp := w.Main.Temp

	tempDictionary := map[float64]string{
		30:	"Hot AF",
		25:	"Hot",
		20:	"Warm",
		15:	"Pleasant",
		10:	"Kinda hot; Kinda warm",
		0:	"Drafty",
		-10:"Chilly",
		-20:"Cold",
		-30:"Freezing",
		-40:"Crazy Freezing!",
		-60:"SUPER DUPER FREEZING!!!",
		-273:"This is so cold this is probably wrong",
	}


	for currentCutoffTemp, response := range tempDictionary{
		if temp > currentCutoffTemp{
			feeling = response
			break
		}
	}

	weather := "Not sure what the weather is"
	if w.Weather != nil {
		weather = w.Weather[0].Description
	}

	returnString := fmt.Sprintf("GetWeather in %s\n%s | %s | %.1f°C", w.Name, feeling, strings.Title(weather), temp)

	return returnString, nil
}