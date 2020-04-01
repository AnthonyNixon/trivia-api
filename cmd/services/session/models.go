package session

import "time"

type Session struct {
	Id          int       `json:"id"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	StartTime   time.Time `json:"startTime"`
	Description string    `json:"description"`
}
