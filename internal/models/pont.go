package models

type Point struct {
	PointID int    `json:"point_id" db:"point_id"`
	Name    string `json:"name" db:"name"`
	Alias   string `json:"alias" db:"alias"`
	Pass    string `json:"pass" db:"pass"`
}
