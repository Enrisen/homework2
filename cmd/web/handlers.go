package main

import (
	"net/http"
	"time"

	"github.com/cohune-cabbage/di/internal/data"
	"github.com/cohune-cabbage/di/internal/validator"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := NewTemplateData()
	data.Title = "Welcome"
	data.HeaderText = "We are here to help"
	err := app.render(w, http.StatusOK, "home.tmpl", data)
	if err != nil {
		app.logger.Error("failed to render home page", "template", "home.tmpl", "error", err, "url", r.URL.Path, "method", r.Method)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

func (app *application) createFeedback(w http.ResponseWriter, r *http.Request) {
	//A. parse the form data
	err := r.ParseForm()
	if err != nil {
		app.logger.Error("failed to parse form", "error", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	name := r.PostForm.Get("name")
	email := r.PostForm.Get("email")
	subject := r.PostForm.Get("subject")
	message := r.PostForm.Get("message")

	//C. Create a Feedback instance using the form data
	//   Remember the Insert method expects a *Feedback
	feedback := &data.Feedback{
		Fullname: name,
		Email:    email,
		Subject:  subject,
		Message:  message,
	}

	v := validator.NewValidator()
	data.ValidateFeedback(v, feedback)

	if !v.ValidData() {
		data := NewTemplateData()
		data.Title = "Welcome"
		data.HeaderText = "We are here to help"
		data.FormErrors = v.Errors
		data.FormData = map[string]string{
			"name":    name,
			"email":   email,
			"subject": subject,
			"message": message,
		}

		err := app.render(w, http.StatusUnprocessableEntity, "home.tmpl", data)
		if err != nil {
			app.logger.Error("failed to render home page", "template", "home.tmpl", "error",
				err, "url", r.URL.Path, "method", r.Method)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		return
	}

	err = app.feedback.Insert(feedback)
	if err != nil {
		app.logger.Error("failed to insert feedback", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/feedback/success", http.StatusSeeOther)
}

func (app *application) feedbackSuccess(w http.ResponseWriter, r *http.Request) {
	data := NewTemplateData()
	data.Title = "Feedback Submitted"
	data.HeaderText = "Thank You for Your Feedback!"
	err := app.render(w, http.StatusOK, "feedback_success.tmpl", data)
	if err != nil {
		app.logger.Error("failed to render feedback success page", "template", "feedback_success.tmpl", "error", err, "url", r.URL.Path, "method", r.Method)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (app *application) feedbackForm(w http.ResponseWriter, r *http.Request) {
	data := NewTemplateData()
	data.Title = "Feedback"
	data.HeaderText = "Share Your Thoughts"
	app.render(w, http.StatusOK, "feedback.tmpl", data)
}

func (app *application) createJournal(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.logger.Error("failed to parse form", "error", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	dateStr := r.PostForm.Get("date")

	var date time.Time
	if dateStr != "" {
		date, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			v := validator.NewValidator()
			v.AddError("date", "Invalid date format")
			data := NewTemplateData()
			data.Title = "Journal"
			data.HeaderText = "New Journal Entry"
			data.FormErrors = v.Errors
			data.FormData = map[string]string{
				"title":   title,
				"content": content,
				"date":    dateStr,
			}

			err := app.render(w, http.StatusUnprocessableEntity, "journal.tmpl", data)
			if err != nil {
				app.logger.Error("failed to render journal page", "template", "journal.tmpl", "error", err, "url", r.URL.Path, "method", r.Method)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			return
		}
	}

	journal := &data.Journal{
		Title:   title,
		Content: content,
		Date:    date,
	}

	v := validator.NewValidator()
	data.ValidateJournal(v, journal)

	if !v.ValidData() {
		data := NewTemplateData()
		data.Title = "Journal"
		data.HeaderText = "New Journal Entry"
		data.FormErrors = v.Errors
		data.FormData = map[string]string{
			"title":   title,
			"content": content,
			"date":    dateStr,
		}

		err := app.render(w, http.StatusUnprocessableEntity, "journal.tmpl", data)
		if err != nil {
			app.logger.Error("failed to render journal page", "template", "journal.tmpl", "error", err, "url", r.URL.Path, "method", r.Method)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		return
	}

	err = app.journal.Insert(journal)
	if err != nil {
		app.logger.Error("failed to insert journal entry", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/journal/success", http.StatusSeeOther)
}

func (app *application) journalForm(w http.ResponseWriter, r *http.Request) {
	data := NewTemplateData()
	data.Title = "Journal"
	data.HeaderText = "New Journal Entry"
	app.render(w, http.StatusOK, "journal.tmpl", data)
}
