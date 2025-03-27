package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /feedback", app.feedbackForm)
	mux.HandleFunc("POST /feedback/new", app.createFeedback)
	mux.HandleFunc("GET /feedback/success", app.feedbackSuccess)
	mux.HandleFunc("GET /journal", app.journalForm)
	mux.HandleFunc("POST /journal/new", app.createJournal)

	// Journal routes
	mux.HandleFunc("GET /journal", app.journalForm)
	mux.HandleFunc("POST /journal/new", app.createJournal)
	mux.HandleFunc("GET /journal/success", app.journalSuccess)

	// Todo routes
	mux.HandleFunc("GET /todos", app.todoListHandler)
	mux.HandleFunc("POST /todos", app.todoCreateHandler)
	mux.HandleFunc("POST /todos/{id}/complete", app.todoCompleteHandler)
	mux.HandleFunc("POST /todos/{id}/delete", app.todoDeleteHandler)

	return app.loggingMiddleware(mux)
}
