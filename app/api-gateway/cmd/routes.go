package main

func (app *Application) initRoutes() {
	e.GET("/", home)
	e.GET("/login", getLogin)
	e.POST("/login", postLogin)
	e.POST("/logout", postLogout)
	e.GET("/register", getRegister)
	e.POST("/register", postRegister)
}
