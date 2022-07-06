package repo

import (
	"calendar-server/internal/domain/event"
	"calendar-server/internal/errors"
	"github.com/google/uuid"
	"log"
	"sync"
	"time"
)

type CalendarStorage struct {
	mx     *sync.RWMutex
	events map[string]*event.Event
}

func NewCalendarStorage() *CalendarStorage {
	return &CalendarStorage{
		mx:     &sync.RWMutex{},
		events: make(map[string]*event.Event),
	}
}

func (c *CalendarStorage) GetEvent(id string) (*event.Event, bool) {
	c.mx.RLock()
	defer c.mx.RUnlock()

	v, ok := c.events[id]
	return v, ok
}

func (c *CalendarStorage) AddEvent(e *event.CreateEventDTO) (*event.Event, error) {
	c.mx.Lock()
	defer c.mx.Unlock()

	event := &event.Event{}

	event.UserId = e.UserId
	event.Name = e.Name
	event.Date = e.Date
	event.ID = uuid.New().String()
	log.Println("event has been added with ID:", event.ID)

	c.events[event.ID] = event

	return event, nil
}

func (c *CalendarStorage) UpdateEvent(e *event.UpdateEventDTO) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	if _, ok := c.events[e.ID]; !ok {
		return errors.EventNotFound
	}

	if !e.Date.IsZero() {
		c.events[e.ID].Date = e.Date
	}

	if e.Name != "" {
		c.events[e.ID].Name = e.Name
	}

	return nil
}

func (c *CalendarStorage) DeleteEvent(id string) {
	c.mx.Lock()
	defer c.mx.Unlock()

	delete(c.events, id)
}

func (c *CalendarStorage) GetEventByPeriod(userId string, start, end time.Time) []*event.Event {
	c.mx.RLock()
	defer c.mx.RUnlock()

	var events []*event.Event

	for _, e := range c.events {
		if e.UserId == userId && (e.Date.Time == start || e.Date.After(start)) && e.Date.Before(end) {
			events = append(events, e)
		}
	}

	return events
}
