package main

import (
	"net/http"

	"github.com/Caps1d/task-manager-cloud-app/auth/pb"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Handlers

func home(c echo.Context) error {
	// 1. If user is not logged in -> show hero landing page or whatever its called, include a login form with register option
	// -> Handle auth with gRPC requests to Auth service
	// 2.0 If user is logged in -> show team space

	e.Use(middleware.Timeout())
	// c.Request().Context() returns context.Context
	r, err := authClient.Login(c.Request().Context(), &pb.LoginRequest{Email: "abagalamaga@crazy.ua", Password: "crazymane"})
	if err != nil {
		app.errorLog.Printf("Failed at authClient login request: %v", err)
		return err
	}

	app.infoLog.Printf("LoginResponse data: status = %v, token = %v", r.Status, r.Token)

	return c.String(http.StatusOK, "Hello, World!")
}

func getLogin(c echo.Context) error {
	return c.String(http.StatusOK, "User endpoint reached")
}
