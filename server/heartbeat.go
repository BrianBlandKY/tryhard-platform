package socket

import (
    "time"
)

type heartbeat struct {
	latencyTime	time.Time
	latency		int
	ticker		*time.Ticker
}
