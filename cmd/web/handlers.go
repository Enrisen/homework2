package main

import (
	"errors"
	"net/http"
<<<<<<< HEAD
=======
	"strconv"
>>>>>>> 4be8292 (Added the todo page and styled it. Still need better styling.)
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

<<<<<<< HEAD
func (app *application) createJournal(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.logger.Error("failed to parse form", "error", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
=======
// --- Journal Handlers ---

func (app *application) journalForm(w http.ResponseWriter, r *http.Request) {
	data := NewTemplateData()
	data.Title = "Journal"
	data.HeaderText = "Write in Your Journal"
	err := app.render(w, http.StatusOK, "journal.tmpl", data)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) createJournal(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

>>>>>>> 4be8292 (Added the todo page and styled it. Still need better styling.)
	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	dateStr := r.PostForm.Get("date")

<<<<<<< HEAD
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
=======
	// Parse the date string into a time.Time
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
>>>>>>> 4be8292 (Added the todo page and styled it. Still need better styling.)
	}

	journal := &data.Journal{
		Title:   title,
		Content: content,
		Date:    date,
	}

	v := validator.NewValidator()
	data.ValidateJournal(v, journal)

	if !v.ValidData() {
<<<<<<< HEAD
		data := NewTemplateData()
		data.Title = "Journal"
		data.HeaderText = "New Journal Entry"
=======
		// If validation fails, re-render the form with errors
		data := NewTemplateData()
		data.Title = "Journal"
		data.HeaderText = "Write in Your Journal"
>>>>>>> 4be8292 (Added the todo page and styled it. Still need better styling.)
		data.FormErrors = v.Errors
		data.FormData = map[string]string{
			"title":   title,
			"content": content,
			"date":    dateStr,
		}

<<<<<<< HEAD
		err := app.render(w, http.StatusUnprocessableEntity, "journal.tmpl", data)
		if err != nil {
			app.logger.Error("failed to render journal page", "template", "journal.tmpl", "error", err, "url", r.URL.Path, "method", r.Method)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
=======
		err = app.render(w, http.StatusUnprocessableEntity, "journal.tmpl", data)
		if err != nil {
			app.serverError(w, r, err)
>>>>>>> 4be8292 (Added the todo page and styled it. Still need better styling.)
		}
		return
	}

	err = app.journal.Insert(journal)
	if err != nil {
<<<<<<< HEAD
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
=======
		app.serverError(w, r, err)
		return
	}

	// Redirect to success page
	http.Redirect(w, r, "/journal/success", http.StatusSeeOther)
}

func (app *application) journalSuccess(w http.ResponseWriter, r *http.Request) {
	data := NewTemplateData()
	data.Title = "Journal Entry Submitted"
	data.HeaderText = "Journal Entry Submitted"
	err := app.render(w, http.StatusOK, "journal_success.tmpl", data)
	if err != nil {
		app.serverError(w, r, err)
	}
}

// --- Todo Handlers ---

// todoListHandler displays the list of todo items.
func (app *application) todoListHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := app.todos.GetAll()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := NewTemplateData()
	data.Title = "Todo List"
	data.HeaderText = "Manage Your Tasks"
	data.Todos = todos

	err = app.render(w, http.StatusOK, "todo.tmpl", data)
	if err != nil {
		app.serverError(w, r, err)
	}
}

// todoCreateHandler handles the creation of a new todo item.
func (app *application) todoCreateHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	task := r.PostForm.Get("task")

	todo := &data.Todo{
		Task:      task,
		Completed: false, // New tasks are not completed by default
	}

	v := validator.NewValidator()
	// Basic validation: task cannot be empty
	v.Check(task != "", "task", "Task cannot be empty")
	// Add more validation rules if needed, e.g., max length

	if !v.ValidData() {
		// If validation fails, re-render the form with errors
		todos, err := app.todos.GetAll() // Fetch existing todos to display
		if err != nil {
			app.serverError(w, r, err)
			return
		}
		data := NewTemplateData()
		data.Title = "Todo List"
		data.HeaderText = "Manage Your Tasks"
		data.Todos = todos
		data.FormErrors = v.Errors
		data.FormData = map[string]string{"task": task}

		err = app.render(w, http.StatusUnprocessableEntity, "todo.tmpl", data)
		if err != nil {
			app.serverError(w, r, err)
		}
		return
	}

	err = app.todos.Insert(todo)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Redirect back to the todo list page
	http.Redirect(w, r, "/todos", http.StatusSeeOther)
}

// todoCompleteHandler marks a todo item as complete.
func (app *application) todoCompleteHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id < 1 {
		app.notFound(w, r)
		return
	}

	todo, err := app.todos.Get(id)
	if err != nil {
		if errors.Is(err, data.ErrRecordNotFound) {
			app.notFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	// Mark as complete and update
	todo.Completed = true
	err = app.todos.Update(todo)
	if err != nil {
		if errors.Is(err, data.ErrEditConflict) {
			// Handle potential edit conflict if necessary
			app.clientError(w, http.StatusConflict) // Or redirect with a message
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	// Redirect back to the todo list page
	http.Redirect(w, r, "/todos", http.StatusSeeOther)
}

// todoDeleteHandler deletes a todo item.
func (app *application) todoDeleteHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id < 1 {
		app.notFound(w, r)
		return
	}

	err = app.todos.Delete(id)
	if err != nil {
		if errors.Is(err, data.ErrRecordNotFound) {
			app.notFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	// Redirect back to the todo list page
	http.Redirect(w, r, "/todos", http.StatusSeeOther)
>>>>>>> 4be8292 (Added the todo page and styled it. Still need better styling.)
}
