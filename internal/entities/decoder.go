package entities

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type CustomTime struct {
	time.Time
}

const customTimeLayout = "2006-01-02"

func (ct *CustomTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	parse, err := time.Parse(customTimeLayout, s)
	if err != nil {
		return fmt.Errorf("invalid date format, expected YYYY-MM-DD")
	}

	ct.Time = parse
	return nil
}

func (ct *CustomTime) Scan(value interface{}) error {
	if v, ok := value.(time.Time); ok {
		*ct = CustomTime{Time: v}
		return nil
	}
	return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type CustomTime", value)
}

func (ct CustomTime) Value() (driver.Value, error) {
	return ct.Time, nil
}
