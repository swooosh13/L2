package event

import (
	"time"
)

var (
	LayoutFullDate = `"` + "2006-01-02" + `"`
	LayoutDateForDay = "2006-01-02"
	LayoutDateForMonth = "2006-01"
)


type CreateEventDTO struct {
	UserId string `json:"user_id" validate:"required"`
	Name string `json:"name" validate:"required"`
	Date Date `json:"date" validate:"required"`
}

type UpdateEventDTO struct {
	ID string `json:"id" validate:"required"`
	Name string `json:"name"`
	Date Date `json:"date"`
}

type DeleteEventDTO struct {
	ID string `json:"id" validate:"required"`
}

type Date struct {
	time.Time
}

func (t *Date) UnmarshalJSON(date []byte) error {
	if string(date) == "" || string(date) == `""` || string(date) == "null" {
		*t = Date{time.Now()}
		return nil
	}
	tm, err := time.Parse(LayoutFullDate, string(date))
	*t = Date{tm}
	return err
}