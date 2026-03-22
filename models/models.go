package models

import "time"

type Availability struct {
	CoachID   int    `json:"coach_id"`
	Day       string `json:"day"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type Booking struct {
	UserID  int       `json:"user_id"`
	CoachID int       `json:"coach_id"`
	Slot    time.Time `json:"datetime"`
}