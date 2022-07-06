package event

type Event struct {
	ID     string `json:"id" validate:"required"`
	UserId string `json:"user_id" validate:"required"`
	Name   string `json:"name" validate:"required"`
	Date   Date `json:"date" validate:"required"`
}