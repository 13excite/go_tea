package api

// FakeAPI implemets test API
type FakeAPI struct {
	name string
}

func New() *FakeAPI {
	return &FakeAPI{
		name: "my-test-api",
	}
}

// WeatherResponse is a map in date=>temp format
type WeatherResponse map[string]string

// GetWetherByCountry returns map[date]temp for the given city
func (api *FakeAPI) GetWetherByCity(city string) WeatherResponse {
	return WeatherResponse{
		"today":    "+18",
		"tomorrow": "+22",
		"1.08":     "+20",
		"2.08":     "+25",
		"3.08":     "+24",
		"4.08":     "+26",
		"5.08":     "+18",
		"6.08":     "+21",
	}
}

// GetAvailableCities returns list of cities
func (api *FakeAPI) GetAvailableCities() []string {
	return []string{
		"Berlin",
		"Munchen",
		"Frankfurt",
		"Leipzig",
		"Longon",
		"Paris",
		"Liverpool",
		"Koln",
		"Lion",
		"Flensburg",
		"Bordo",
		"Erfurt",
		"Dresden",
	}
}
