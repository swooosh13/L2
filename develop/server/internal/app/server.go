package app

import (
	"calendar-server/internal/handlers"
	"calendar-server/internal/repo"
	"log"
	"net"
	"net/http"
)

type Server struct {
	storage *repo.CalendarStorage
	router  *http.ServeMux
}

func NewServer() *Server {
	return &Server{
		storage: repo.NewCalendarStorage(),
		router:  http.NewServeMux(),
	}
}

func (s *Server) Run(host, port string) {
	srv := &http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: s.router,
	}

	s.router.HandleFunc("/create_event", handlers.LogRequest(s.CreateEvent))
	s.router.HandleFunc("/update_event", handlers.LogRequest(s.UpdateEvent))
	s.router.HandleFunc("/delete_event", handlers.LogRequest(s.DeleteEvent))
	s.router.HandleFunc("/events_for_day", handlers.LogRequest(s.GetEventForDay))
	s.router.HandleFunc("/events_for_week", handlers.LogRequest(s.GetEventForWeek))
	s.router.HandleFunc("/events_for_month", handlers.LogRequest(s.GetEventForMonth))

	log.Println("server has been started")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
