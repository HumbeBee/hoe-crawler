package definitions

import (
	"database/sql/driver"
	"fmt"
)

type HoeStatus string

const (
	HoeStatusActive   HoeStatus = "active"
	HoeStatusInactive HoeStatus = "inactive"
	HoeStatusUnknown  HoeStatus = "unknown"
)

// For database serialization
func (s HoeStatus) Value() (driver.Value, error) {
	return string(s), nil
}

// For database deserialization
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
