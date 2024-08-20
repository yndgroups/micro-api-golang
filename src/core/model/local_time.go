package model

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

const localDateTimeFormat string = "2006-01-02 15:04:05"

type LocalTime time.Time

func (t LocalTime) MarshalJSON() ([]byte, error) {
	format := fmt.Sprintf(`"%s"`, time.Time(t).Format(localDateTimeFormat))
	if strings.Contains(format, "0001") {
		return []byte("null"), nil
	}
	return []byte(format), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if time.Time(t).UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return time.Time(t), nil
}
