package entities

import "time"

type StreamingStatus struct {
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
}
