package models

import "time"

type GPS struct {
	GpsID      int       `json:"gps_id" db:"gps_id"`
	PointID    int       `json:"point_id" db:"point_id"`
	PointGpsID int       `json:"point_gps_id" db:"point_gps_id"`
	Lat        float64   `json:"lat" db:"lat"`
	Lon        float64   `json:"lon" db:"lon"`
	Speed      float64   `json:"speed" db:"speed"`
	Time       time.Time `json:"time" db:"time"`
}
