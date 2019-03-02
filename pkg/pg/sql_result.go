package pg

import (
	"fmt"
	"time"
)

// SQLResult tracks sql response and time taken
type SQLResult struct {
	Timestamp QueryStart  `json:"timestamp"`
	Message   string      `json:"message"`
	TimeTaken QueryTime   `json:"time_taken"`
	Status    QueryStatus `json:"status"`
}

// QueryStatus represents if the query run was a success or failure
type QueryStatus string

// QueryStart is the timestamp the query started
type QueryStart time.Time

// MarshalJSON to format the QueryStart timestamp
func (t QueryStart) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("15:04:05"))
	return []byte(stamp), nil
}

// QueryTime represents the amount of time the query took
type QueryTime float64

// MarshalJSON to format the QueryTime
func (t QueryTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%.3fms\"", t)
	return []byte(stamp), nil
}

const (
	success QueryStatus = "success"
	failure QueryStatus = "failed"
)
