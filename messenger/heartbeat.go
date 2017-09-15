package messenger

import (
	"time"
)

// Heartbeat Model
type Heartbeat struct {
	NodeStartTime       time.Time // Node Start Time
	NodeEndTime         time.Time // Node End Time
	ServiceReceivedTime time.Time // Service Received Time
	UploadLatency       float64   // Upload Latency MS
	DownloadLatency     float64   // Download Latency MS
}
