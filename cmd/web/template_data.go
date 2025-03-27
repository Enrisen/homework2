package main

import "github.com/cohune-cabbage/di/internal/data"

type TemplateData struct {
	Title      string
	HeaderText string
	FormErrors map[string]string
	FormData   map[string]string
	Todos      []*data.Todo // Add Todos field
	CSRFToken  string       // Add CSRF token field
}

func NewTemplateData() *TemplateData {
	return &TemplateData{
		Title:      "Default Title",
		HeaderText: "Default HeaderText",
		FormErrors: map[string]string{},
		FormData:   map[string]string{},
		Todos:      []*data.Todo{}, // Initialize Todos
	}
}
