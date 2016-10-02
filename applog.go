package ouchapp

import (
	"time"
)

type AppLog struct {
	ID          string    `json:"id"`
	UserID      string    `json:"userid"`
	LogDateTime time.Time `json:"logdatetime"`
	LogLevel    int       `json:"loglevel"`
	LogType     int       `json:"logtype"`
	LogMessage  string    `json:"logmessage"`
}
