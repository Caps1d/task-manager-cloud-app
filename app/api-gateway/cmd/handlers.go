package main

import (
	"net/http"

	"github.com/Caps1d/task-manager-cloud-app/auth/pb"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type LoginForm struct {
	Email    string
	Password string
}

type RegistrationForm struct {
	Email    string
	Password string
	Username string
}

// Handlers
func home(c echo.Context) error {
	// 1. If user is not logged in -> show hero landing page or whatever its called, include a login form with register option
	// -> Handle auth with gRPC requests to Auth service
	// 2.0 If user is logged in -> show team space

	return c.String(http.StatusOK, "Hello, World!")
}

func getLogin(c echo.Context) error {
	return c.String(http.StatusOK, "User endpoint reached")
}

func postLogin(c echo.Context) error {
	e.Use(middleware.Timeout())
	// c.Request().Context() returns context.Context
	var form LoginForm

	err := c.Bind(&form)
	if err != nil {
		app.errorLog.Printf("Failed to bind login request body: %v", err)
		return c.String(http.StatusBadRequest, "bad request")
	}

	app.infoLog.Printf("Post login request email: %v", form.Email)

	email := form.Email
	password := form.Password

	r, err := authClient.Login(c.Request().Context(), &pb.LoginRequest{Email: email, Password: password})
	if err != nil {
		app.errorLog.Printf("Failed at authClient login request: %v", err)
		return err
	}

	app.infoLog.Printf("LoginResponse data: status = %v, token = %v", r.Status, r.Token)

	return c.String(http.StatusOK, "Login successful")
}

func getRegister(c echo.Context) error {
	return c.String(http.StatusOK, "Register endpoint reached")
}

func postRegister(c echo.Context) error {
	var form RegistrationForm

	e.Use(middleware.Timeout())

	err := c.Bind(&form)
	if err != nil {
		app.errorLog.Printf("Failed to bind login request body: %v", err)
		return c.String(http.StatusBadRequest, "bad request")
	}

	email := form.Email
	password := form.Password
	username := form.Username

	app.infoLog.Printf("Echo server postRegister email = %v", email)

	r, err := authClient.Register(c.Request().Context(), &pb.RegisterRequest{Email: email, Password: password, Username: username})
	if err != nil {
		app.errorLog.Printf("Failed at authClient register request: %v", err, r.Error)
		return err
	}

	app.infoLog.Printf("RegisterResponse data: status = %v", r.Status)

	return c.String(http.StatusOK, "Register endpoint reached")
}
