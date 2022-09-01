package model

import "time"

type Run struct {
	Id                int64        `json:"id" db:"id"`
	LocationName      string       `json:"location_name" db:"location_name"`
	Date              time.Time    `json:"date" db:"date"`
	Duration          int          `json:"duration" db:"duration"`
	Distance          int          `json:"distance" db:"distance"`
	TotalSteps        int          `json:"total_steps" db:"total_steps"`
	Routes            []Coordinate `json:"route"`
	FormattedDate     string       `json:"formatted_date"`
	FormattedDuration string       `json:"formatted_duration"`
}

type Coordinate struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}
