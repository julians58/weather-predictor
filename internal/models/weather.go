package models

type Weather struct {
	Day       int    `json:"day" gorm:"primaryKey"`
	Condition string `json:"condition"`
}

type WeatherStatistics struct {
	ID             uint   `gorm:"primaryKey"`
	DroughtPeriods int    `json:"drought_periods"`
	RainPeriods    int    `json:"rain_periods"`
	OptimalPeriods int    `json:"optimal_periods"`
	RainPeakDays   string `json:"rain_peak_days"` // Almacenar como JSON en formato de string
}
