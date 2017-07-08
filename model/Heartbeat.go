package model

import (
	"time"
)

// Heartbeat Model
type Heartbeat struct {
	LatencyTime    time.Time    // timestamp
	LatencySeconds int          // seconds
	Interval       *time.Ticker // interval between heartbeat requests
}
