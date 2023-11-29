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

	switch city {
	case "berlin":
		return WeatherResponse{
			"today":    "+3",
			"tomorrow": "+4",
			"1.12":     "+3",
		}
	case "munchen":
		return WeatherResponse{
			"today":    "0",
			"tomorrow": "-1",
			"1.12":     "+2",
		}
	case "frankfurt":
		return WeatherResponse{
			"today":    "+1",
			"tomorrow": "0",
			"1.12":     "+4",
		}
	case "leipzig":
		return WeatherResponse{
			"today":    "+2",
			"tomorrow": "+4",
			"1.12":     "+2",
		}
	case "longon":
		return WeatherResponse{
			"today":    "+8",
			"tomorrow": "+10",
			"1.12":     "+11",
		}
	default:
		return WeatherResponse{
			"today":    "+18",
			"tomorrow": "+22",
			"1.12":     "+20",
		}
	}
}
