package entities

type WeatherPayload struct {
	Data struct {
		Request []struct {
			Type  string `json:"type"`
			Query string `json:"query"`
		} `json:"request"`
		Weather []struct {
			Date      string `json:"date"`
			Astronomy []struct {
				Sunrise          string `json:"sunrise"`
				Sunset           string `json:"sunset"`
				Moonrise         string `json:"moonrise"`
				Moonset          string `json:"moonset"`
				MoonPhase        string `json:"moon_phase"`
				MoonIllumination string `json:"moon_illumination"`
			} `json:"astronomy"`
			MaxtempC    string `json:"maxtempC"`
			MaxtempF    string `json:"maxtempF"`
			MintempC    string `json:"mintempC"`
			MintempF    string `json:"mintempF"`
			TotalSnowCm string `json:"totalSnow_cm"`
			SunHour     string `json:"sunHour"`
			UvIndex     string `json:"uvIndex"`
			Hourly      []struct {
				Time           string `json:"time"`
				TempC          string `json:"tempC"`
				TempF          string `json:"tempF"`
				WindspeedMiles string `json:"windspeedMiles"`
				WindspeedKmph  string `json:"windspeedKmph"`
				WinddirDegree  string `json:"winddirDegree"`
				Winddir16Point string `json:"winddir16Point"`
				WeatherCode    string `json:"weatherCode"`
				WeatherIconURL []struct {
					Value string `json:"value"`
				} `json:"weatherIconUrl"`
				WeatherDesc []struct {
					Value string `json:"value"`
				} `json:"weatherDesc"`
				PrecipMM      string `json:"precipMM"`
				Humidity      string `json:"humidity"`
				Visibility    string `json:"visibility"`
				Pressure      string `json:"pressure"`
				Cloudcover    string `json:"cloudcover"`
				HeatIndexC    string `json:"HeatIndexC"`
				HeatIndexF    string `json:"HeatIndexF"`
				DewPointC     string `json:"DewPointC"`
				DewPointF     string `json:"DewPointF"`
				WindChillC    string `json:"WindChillC"`
				WindChillF    string `json:"WindChillF"`
				WindGustMiles string `json:"WindGustMiles"`
				WindGustKmph  string `json:"WindGustKmph"`
				FeelsLikeC    string `json:"FeelsLikeC"`
				FeelsLikeF    string `json:"FeelsLikeF"`
			} `json:"hourly"`
		} `json:"weather"`
	} `json:"data"`
}
