package agent

import "time"

type StreamLog struct {
	Level     string    `json:"@level"`
	Message   string    `json:"@message"`
	Module    string    `json:"@module"`
	Timestamp time.Time `json:"@timestamp"`
	Check     string    `json:"check"`
}
