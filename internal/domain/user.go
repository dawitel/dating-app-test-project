package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
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

// Implement the Scanner interface for Location to scan data from the database
func (l *Location) Scan(value interface{}) error {
	// Try to convert the value to a byte slice
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to convert database value to Location")
	}

	// Unmarshal the JSON data into the Location struct
	if err := json.Unmarshal(bytes, l); err != nil {
		return fmt.Errorf("failed to unmarshal Location: %v", err)
	}

	return nil
}

// Implement the Valuer interface for Location to store the Location in the database
func (l Location) Value() (driver.Value, error) {
	// Marshal the Location struct into a JSON format for storage
	return json.Marshal(l)
}