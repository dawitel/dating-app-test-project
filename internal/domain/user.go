package domain

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/lib/pq"
)

// Location represents a geographical location with latitude and longitude
type Location struct {
	Latitude  float64 `gorm:"column:latitude" json:"latitude"`
	Longitude float64 `gorm:"column:longitude" json:"longitude"`
}

// User represets the users in the db.
type User struct {
	UserID      string         `json:"user_id" gorm:"primaryKey"`
	Name        string         `json:"name"`
	Password    string         `json:"-"`
	Age         int            `json:"age"`
	Gender      string         `json:"gender"`
	Location    Location       `json:"location"`
	Interests   pq.StringArray `json:"interests" gorm:"type:text[]"`
	Preferences Preferences    `json:"preferences" gorm:"type:jsonb"`
	LastActive  time.Time      `json:"last_active"`
	Score       int            `json:"-"` // Ignore this field when persisting to DB
}

type Preferences struct {
	MinAge      int    `json:"min_age"`
	MaxAge      int    `json:"max_age"`
	Gender      string `json:"preferred_gender"`
	MaxDistance int    `json:"max_distance"` // Distance in kilometers
}

// Implement the Scanner interface for Preferences to deserialize JSONB from DB
func (p *Preferences) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to convert database value to Preferences")
	}
	return json.Unmarshal(bytes, p)
}

// Implement the Valuer interface for Preferences to serialize to JSONB for DB
func (p Preferences) Value() (driver.Value, error) {
	return json.Marshal(p)
}
