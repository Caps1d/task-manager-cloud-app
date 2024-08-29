package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	authpb "github.com/Caps1d/task-manager-cloud-app/auth/pb"
	userpb "github.com/Caps1d/task-manager-cloud-app/user/pb"
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
	Name     string
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
		r, err := authClient.IsAuthenticated(context.Background(), &authpb.IsAuthenticatedRequest{SessionID: cookie.Value})
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

	r, err := authClient.Login(c.Request().Context(), &authpb.LoginRequest{Email: email, Password: password})
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

	r, err := authClient.Logout(context.Background(), &authpb.LogoutRequest{SessionID: cookie.Value})
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

	err := c.Bind(&form)
	if err != nil {
		app.errorLog.Printf("Failed to bind register request body: %v", err)
		return c.String(http.StatusBadRequest, "bad request")
	}

	email := form.Email
	password := form.Password
	username := form.Username
	name := form.Name

	app.infoLog.Printf("Echo server postRegister email = %v", email)

	a, err := authClient.Register(c.Request().Context(), &authpb.RegisterRequest{Email: email, Password: password, Username: username})
	if err != nil {
		app.errorLog.Printf("API: error at auth register request: %v", err)
		return err
	}

	app.infoLog.Printf("RegisterResponse data from auth service: success = %v", a.Success)

	u, err := userClient.Register(c.Request().Context(), &userpb.RegisterRequest{Email: email, Name: name, Username: username})
	if err != nil {
		app.errorLog.Printf("API: error at user register request: %v", err)
		return err
	}

	app.infoLog.Printf("RegisterResponse data from user service: success = %v", u.Success)

	return c.String(http.StatusOK, "Register endpoint reached")
}

type User struct {
	Id       int32
	Name     string
	Email    string
	Username string
	Role     string
	TeamId   int32
}

func getUser(c echo.Context) error {
	param := c.QueryParam("id")
	uid, err := strconv.Atoi(param)
	if err != nil {
		app.errorLog.Printf("API: failed to parse id from request, %v", err)
		return err
	}

	r, err := userClient.GetUser(c.Request().Context(), &userpb.GetUserRequest{Id: int32(uid)})
	if err != nil {
		app.errorLog.Printf("API: error at get user request: %v", err)
		return err
	}

	data := r.Data

	u := &User{
		Id:       data.Id,
		Name:     data.Name,
		Email:    data.Email,
		Username: data.Username,
	}

	app.infoLog.Printf("API: got user with username: %v", u.Username)

	return c.JSON(http.StatusOK, u)
}
