package main

import (
	"fmt"
	"sync"
	"time"
)

// Store — place to store all user events.
type Store struct {
	mu     *sync.Mutex
	events map[int][]Event
}

// Create — adds new instance of the event to the storage.
func (s *Store) Create(e *Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if events, ok := s.events[e.UserID]; ok {
		for _, event := range events {
			if event.EventID == e.EventID {
				return fmt.Errorf("event with such id (%v) already present for this user (%v);", e.EventID, e.UserID)
			}
		}
	}
	s.events[e.UserID] = append(s.events[e.UserID], *e)
	return nil
}

// Update — updates the existing instance of the event in the storage.
func (s *Store) Update(e *Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	index := -1
	events := make([]Event, 0)
	ok := false
	if events, ok = s.events[e.UserID]; !ok {
		return fmt.Errorf("user with such id (%v) does not exists;", e.UserID)
	}
	for idx, event := range events {
		if event.EventID == e.EventID {
			index = idx
			break
		}
	}
	if index == -1 {
		return fmt.Errorf("there is no event with such id (%v) for this user (%v);", e.EventID, e.UserID)
	}
	s.events[e.UserID][index] = *e
	return nil
}

// Delete — deletes event from the storage.
func (s *Store) Delete(e *Event) (*Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	index := -1
	events := make([]Event, 0)
	ok := false
	if events, ok = s.events[e.UserID]; !ok {
		return nil, fmt.Errorf("user with such id (%v) does not exists;", e.UserID)
	}
	for idx, event := range events {
		if event.EventID == e.EventID {
			index = idx
			break
		}
	}
	if index == -1 {
		return nil, fmt.Errorf("there is no event with such id (%v) for this user (%v);", e.EventID, e.UserID)
	}
	eventsLength := len(s.events[e.UserID])
	deletedEvent := s.events[e.UserID][index]
	s.events[e.UserID][index] = s.events[e.UserID][eventsLength-1]
	s.events[e.UserID] = s.events[e.UserID][:eventsLength-1]
	// Could be a good idea to free a slice with such a user id if there are no elements left.
	return &deletedEvent, nil
}

// GetEventsForDay — gets all events for the specific day.
func (s *Store) GetEventsForDay(userID int, date time.Time) ([]Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var result []Event
	events := make([]Event, 0)
	ok := false
	if events, ok = s.events[userID]; !ok {
		return nil, fmt.Errorf("user with such id (%v) does not exists;", userID)
	}
	for _, event := range events {
		if event.Date.Year() == date.Year() && event.Date.Month() == date.Month() && event.Date.Day() == date.Day() {
			result = append(result, event)
		}
	}
	return result, nil
}

// GetEventsForWeek — gets all events for the specific week.
func (s *Store) GetEventsForWeek(userID int, date time.Time) ([]Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var result []Event
	events := make([]Event, 0)
	ok := false
	if events, ok = s.events[userID]; !ok {
		return nil, fmt.Errorf("user with such id (%v) does not exists;", userID)
	}
	for _, event := range events {
		y1, w1 := event.Date.ISOWeek()
		y2, w2 := date.ISOWeek()
		if y1 == y2 && w1 == w2 {
			result = append(result, event)
		}
	}
	return result, nil
}

// GetEventsForMonth — gets all events for the specific month.
func (s *Store) GetEventsForMonth(userID int, date time.Time) ([]Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var result []Event
	events := make([]Event, 0)
	ok := false
	if events, ok = s.events[userID]; !ok {
		return nil, fmt.Errorf("user with such id (%v) does not exists;", userID)
	}
	for _, event := range events {
		if event.Date.Year() == date.Year() && event.Date.Month() == date.Month() {
			result = append(result, event)
		}
	}
	return result, nil
}
