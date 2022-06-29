package main

import "fmt"

// Server - конкретный продукт
type Server struct {
	Type string
	Core int
	Mem  int
}

func NewServer() Computer {
	return &Server{
		Type: ServerType,
		Core: 16,
		Mem:  256,
	}
}

func (s *Server) GetType() string {
	return s.Type
}

func (s *Server) PrintDetails() {
	fmt.Printf("%s Core:[%d], Mem:[%d]\n", s.Type, s.Core, s.Mem)
}
