package app

import (
	"calendar-server/internal/domain/event"
	"calendar-server/internal/errors"
	"calendar-server/internal/handlers"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func (s *Server) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if err := handlers.ValidateMethod(r, http.MethodPost); err != nil {
		JsonResponse(w, true, http.StatusInternalServerError, err.Error())
		return
	}

	var eventDTO *event.CreateEventDTO
	if err := json.NewDecoder(r.Body).Decode(&eventDTO); err != nil {
		log.Println(err)
		JsonResponse(w, true, http.StatusInternalServerError, errors.JsonDecodingError.Error())
		return
	}

	if err := event.Validate(eventDTO); err != nil {
		log.Println(err)
		JsonResponse(w, true, http.StatusBadRequest, err.Error())
		return
	}

	s.storage.AddEvent(eventDTO)
	JsonResponse(w, false, http.StatusOK, "event has been created")
}

func (s *Server) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if err := handlers.ValidateMethod(r, http.MethodPost); err != nil {
		JsonResponse(w, true, http.StatusInternalServerError, err.Error())
		return
	}

	var eventDTO *event.UpdateEventDTO
	if err := json.NewDecoder(r.Body).Decode(&eventDTO); err != nil {
		log.Println(err)
		JsonResponse(w, true, http.StatusInternalServerError, errors.JsonDecodingError.Error())
		return
	}

	if err := event.Validate(eventDTO); err != nil {
		log.Println(err)
		JsonResponse(w, true, http.StatusBadRequest, err.Error())
		return
	}

	if err := s.storage.UpdateEvent(eventDTO); err != nil {
		log.Println(err)
		JsonResponse(w, true, http.StatusServiceUnavailable, err.Error())
		return
	}
	JsonResponse(w, false, http.StatusOK, "event has been successfully updated")
}

func (s *Server) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if err := handlers.ValidateMethod(r, http.MethodPost); err != nil {
		JsonResponse(w, true, http.StatusInternalServerError, err.Error())
		return
	}

	var eventDTO *event.DeleteEventDTO
	if err := json.NewDecoder(r.Body).Decode(&eventDTO); err != nil {
		log.Println(err)
		JsonResponse(w, true, http.StatusInternalServerError, errors.JsonDecodingError.Error())
		return
	}

	if err := event.Validate(eventDTO); err != nil {
		log.Println(err)
		JsonResponse(w, true, http.StatusBadRequest, err.Error())
		return
	}

	s.storage.DeleteEvent(eventDTO.ID)
	JsonResponse(w, false, http.StatusOK, "event has been successfully deleted")
}

func (s *Server) GetEventForDay(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")

	date := r.URL.Query().Get("date")
	dateTime, err := time.Parse(event.LayoutDateForDay, date)
	if err != nil {
		log.Println(err.Error())
		JsonResponse(w, true, http.StatusBadRequest, err.Error())
		return
	}

	events := s.storage.GetEventByPeriod(userId, dateTime, dateTime.AddDate(0, 0, 1))
	JsonResponse(w, false, http.StatusOK, events)
}

func (s *Server) GetEventForWeek(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")

	date := r.URL.Query().Get("date")
	dateTime, err := time.Parse(event.LayoutDateForDay, date)
	if err != nil {
		log.Println(err.Error())
		JsonResponse(w, true, http.StatusBadRequest, err.Error())
		return
	}

	events := s.storage.GetEventByPeriod(userId, dateTime, dateTime.AddDate(0, 0, 7))
	JsonResponse(w, false, http.StatusOK, events)
}

func (s *Server) GetEventForMonth(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")

	date := r.URL.Query().Get("date")
	dateTime, err := time.Parse(event.LayoutDateForMonth, date)
	if err != nil {
		log.Println(err.Error())
		JsonResponse(w, true, http.StatusBadRequest, err.Error())
		return
	}

	events := s.storage.GetEventByPeriod(userId, dateTime, dateTime.AddDate(0, 1, 0))
	JsonResponse(w, false, http.StatusOK, events)

}
