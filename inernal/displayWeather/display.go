package display

import (
	weather "CLIproject/api/wether"

	"github.com/cheynewallace/tabby"
)

func DisplayWeatherData(weather *weather.WeatherData) {
	t := tabby.New()
	if len(weather.Alerts) > 0 {
		t.AddHeader("Weather for", "Description", "Temperature", "Humidity")
		t.AddLine(
			weather.Address, 
			weather.Alerts[0].Description,
			weather.CurrentConditions.Temp, 
			weather.CurrentConditions.Humidity,
		)
	}else {
		t.AddHeader("Weather for", "Temperature", "Humidity")
		t.AddLine(
			weather.Address, 
			weather.CurrentConditions.Temp, 
			weather.CurrentConditions.Humidity,
		)
	}
	t.Print()
}