package response

import (
	"encoding/json"
	"time"
)

const (
	DateFormat     = "2006-01-02"
	DateTimeFormat = "2006-01-02 15:04:05"
)

// DateResponse is a custom type for date response.
type Date time.Time

// MarshalJSON implements the json.Marshaler interface.
func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d).Local().Format(DateFormat))
}

// DateTimeResponse is a custom type for datetime response.
type DateTime time.Time

// MarshalJSON implements the json.Marshaler interface.
func (d DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d).Local().Format(DateTimeFormat))
}
