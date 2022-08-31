package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

const dateLayout = "2006-01-02"

// Create global storage to store all events.
var storage Store = Store{events: make(map[int][]Event), mu: &sync.Mutex{}}

func main() {
	mux := http.NewServeMux()
	// POST routes.
	mux.HandleFunc("/create_event", CreateEventHandler)
	mux.HandleFunc("/update_event", UpdateEventHandler)
	mux.HandleFunc("/delete_event", DeleteEventHandler)
	// GET routes.
	mux.HandleFunc("/events_for_day", EventsForDayHandler)
	mux.HandleFunc("/events_for_week", EventsForWeekHandler)
	mux.HandleFunc("/events_for_month", EventsForMonthHandler)
	// Set Logger middleware.
	wrappedMux := NewLogger(mux)
	// Small function to read port from config.
	port := ":8080"
	func() {
		temp := os.Getenv("PORT")
		if temp != "" {
			port = temp
		}
	}()
	log.Printf("Server is listening for incoming requests at: %v", port)
	log.Fatalln(http.ListenAndServe(port, wrappedMux))
}

func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	var e Event
	if err := e.Decode(r.Body); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := e.Validate(); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := storage.Create(&e); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	resultResponse(w, "Event has been created successfully!", []Event{e}, http.StatusCreated)
	fmt.Println(storage.events)
}

// UpdateEventHandler — handler for the /update_event path.
func UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	var e Event
	if err := e.Decode(r.Body); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := e.Validate(); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := storage.Update(&e); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	resultResponse(w, "Event has been updated successfully!", []Event{e}, http.StatusOK)
	fmt.Println(storage.events)
}

// DeleteEventHandler — handler for the /delete_event path.
func DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	var e Event
	if err := e.Decode(r.Body); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	var deletedEvent *Event
	var err error
	if deletedEvent, err = storage.Delete(&e); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	resultResponse(w, "Event has been deleted successfully!", []Event{*deletedEvent}, http.StatusOK)
}

// EventsForDayHandler — handler for the /events_for_day path.
func EventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	date, err := time.Parse(dateLayout, r.URL.Query().Get("date"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	var events []Event
	if events, err = storage.GetEventsForDay(userID, date); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	resultResponse(w, "Request has been executed successfully!", events, http.StatusOK)
}

// EventsForWeekHandler — handler for the /events_for_week path.
func EventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	date, err := time.Parse(dateLayout, r.URL.Query().Get("date"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	var events []Event
	if events, err = storage.GetEventsForWeek(userID, date); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	resultResponse(w, "Request has been executed successfully!", events, http.StatusOK)
}

// EventsForMonthHandler — handler for the /events_for_month path.
func EventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	date, err := time.Parse(dateLayout, r.URL.Query().Get("date"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	var events []Event
	if events, err = storage.GetEventsForMonth(userID, date); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	resultResponse(w, "Request has been executed successfully!", events, http.StatusOK)
}

// errorResponse — response with error.
func errorResponse(w http.ResponseWriter, e string, status int) {
	errorResponse := struct {
		Error string `json:"error"`
	}{Error: e}
	js, err := json.Marshal(errorResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// resultResponse — response with result.
func resultResponse(w http.ResponseWriter, r string, e []Event, status int) {
	resultResponse := struct {
		Result string  `json:"result"`
		Events []Event `json:"events"`
	}{Result: r, Events: e}
	js, err := json.Marshal(resultResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
