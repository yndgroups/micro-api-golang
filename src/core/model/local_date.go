package model

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

const localDateFormat string = "2006-01-02"

type LocalDate time.Time

func (t LocalDate) MarshalJSON() ([]byte, error) {
	//seconds := t.Unix()
	format := fmt.Sprintf(`"%s"`, time.Time(t).Format(localDateFormat))
	if strings.Contains(format, "0001") {
		return []byte("null"), nil
	}
	return []byte(format), nil
}

func (t LocalDate) Value() (driver.Value, error) {
	var zeroTime time.Time
	if time.Time(t).UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return time.Time(t), nil
}
