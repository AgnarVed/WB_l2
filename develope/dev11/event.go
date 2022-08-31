package main

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Event — model of data to store.
type Event struct {
	UserID      int       `json:"user_id"`
	EventID     int       `json:"event_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

// Decode — decodes data from reader (that is coming in json format) to Event struct.
func (e *Event) Decode(r io.Reader) error {
	err := json.NewDecoder(r).Decode(&e)
	if err != nil {
		return err
	}
	return nil
}

// Validate — validates all important fields in the Event struct.
func (e *Event) Validate() error {
	if e.UserID <= 0 {
		return fmt.Errorf("invalid user id: %v;", e.UserID)
	}
	if e.EventID <= 0 {
		return fmt.Errorf("invalid event id: %v;", e.EventID)
	}
	if e.Title == "" {
		return fmt.Errorf("title cannot be empty;")
	}
	return nil
}
