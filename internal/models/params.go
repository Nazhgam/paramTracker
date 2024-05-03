package models

import "time"

type ParamStr struct {
	ID            int       `json:"string_id" db:"string_id"`
	StringPointID int       `json:"string_point_id" db:"string_point_id"`
	PointID       int       `json:"point_id" db:"point_id"`
	Channel       int       `json:"channel" db:"channel"`
	Value         string    `json:"value" db:"value"`
	Time          time.Time `json:"time" db:"time"`
}

type ParamInt struct {
	ID         int       `json:"int_id" db:"int_id"`
	IntPointID int       `json:"int_point_id" db:"int_point_id"`
	PointID    int       `json:"point_id" db:"point_id"`
	Channel    int       `json:"channel" db:"channel"`
	Value      int       `json:"value" db:"value"`
	Time       time.Time `json:"time" db:"time"`
}
