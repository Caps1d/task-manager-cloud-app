package main

import "net/http"

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("Home endpoint reached")
	w.Write([]byte("Home endpoint reached"))
}

// Task
func (app *Application) taskView(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("task view endpoint reached")
}

func (app *Application) taskCreate(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("task create endpoint reached")
}

func (app *Application) taskCreatePost(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("task create endpoint reached")
}

// User
func (app *Application) userSignUp(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("user signup endpoint reached")
}

func (app *Application) userSignUpPost(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("user signup endpoint reached")
}

func (app *Application) userProfile(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("user profile endpoint reached")
}

func (app *Application) userTasksView(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("user tasks endpoint reached")
}

// Team
func (app *Application) teamView(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("team endpoint reached")
}

func (app *Application) teamSignUp(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("team signup form GET endpoint reached")
}

func (app *Application) teamSingUpPost(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("team signup form POST endpoint reached")
}

func (app *Application) teamTaskAssign(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("team signup form POST endpoint reached")
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
