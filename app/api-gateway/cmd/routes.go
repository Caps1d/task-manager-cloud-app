package main

func (app *Application) initRoutes() {
	e.GET("/", home)
	e.GET("/login", getLogin)
}
