package definitions

import (
	"database/sql/driver"
	"fmt"
)

type HoeStatus string

type ParsedAddress struct {
	Street   string
	District string
}

const (
	HoeStatusActive   HoeStatus = "active"
	HoeStatusInactive HoeStatus = "inactive"
	HoeStatusUnknown  HoeStatus = "unknown"
)

// Using value receiver for Value() as we only need to read the status
// Using pointer receiver for Scan() as we need to modify the status
// This mixed receiver pattern is recommended by GORM
// Docs: https://gorm.io/docs/data_types.html

func (s HoeStatus) Value() (driver.Value, error) {
	return string(s), nil
}

func (s *HoeStatus) Scan(value interface{}) error {
	if value == nil {
		*s = HoeStatusUnknown
		return nil
	}

	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("invalid status value: %v", value)
	}

	*s = HoeStatus(str)
	return nil
}
