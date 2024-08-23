package main

import (
	"context"
	"errors"
	"net/http"
	"time"

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
	// check if user-session cookie exists
	cookie, err := c.Cookie("user-session")
	if err != nil && !errors.Is(err, http.ErrNoCookie) {
		app.infoLog.Printf("%v", err)
	}
	// check if sessionID matches
	if cookie != nil {
		r, err := authClient.IsAuthenticated(context.Background(), &pb.IsAuthenticatedRequest{SessionID: cookie.Value})
		if err != nil {
			app.errorLog.Printf("Cookie authentication check failed, error: %v", err)
		}
		if r.Success {
			app.infoLog.Printf("User logged in, redirecting to home page")
			c.Redirect(http.StatusPermanentRedirect, "/")
			return nil
		}
	}

	return c.String(http.StatusOK, "User endpoint reached")
}

func postLogin(c echo.Context) error {
	// c.Request().Context() returns context.Context
	e.Use(middleware.Timeout())

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

	// write sessionID to cookie
	cookie := new(http.Cookie)
	cookie.Name = "user-session"
	cookie.Value = r.Id
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)

	app.infoLog.Printf("Login Successful")

	// should check if on protected route and only then redirect
	return c.Redirect(http.StatusSeeOther, "/")
	// return c.String(http.StatusOK, "User logged in")
}

func postLogout(c echo.Context) error {
	cookie, err := c.Cookie("user-session")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			app.errorLog.Printf("User must be logged in to logout")
		} else {
			app.errorLog.Printf("%v", err)
		}
	}

	r, err := authClient.Logout(context.Background(), &pb.LogoutRequest{SessionID: cookie.Value})
	if err != nil {
		app.errorLog.Printf("Failed at authClient logout request %v", err)
	}

	if r.Success {
		cookie.Value = ""
		cookie.MaxAge = 0
		app.infoLog.Printf("User logged out")
	}

	return c.Redirect(http.StatusSeeOther, "/")
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
		app.errorLog.Printf("Failed at authClient register request: %v", err, r.Success)
		return err
	}

	app.infoLog.Printf("RegisterResponse data: status = %v", r.Success)

	return c.String(http.StatusOK, "Register endpoint reached")
}
