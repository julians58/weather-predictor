package models

type Weather struct {
	Day       int    `json:"day" gorm:"primaryKey"`
	Condition string `json:"condition"`
}
