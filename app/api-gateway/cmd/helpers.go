package main

import (
	"context"
	"net/http"

	authpb "github.com/Caps1d/task-manager-cloud-app/auth/pb"
	"github.com/labstack/echo/v4"
)

func (app *Application) badRequest(c echo.Context, message string) error {
	return c.JSON(http.StatusBadRequest, map[string]string{"error": message})
}

func (app *Application) serverError(c echo.Context, message string) error {
	return c.JSON(http.StatusInternalServerError, map[string]string{"error": message})
}

func (app *Application) unauthorized(c echo.Context, message string) error {
	return c.JSON(http.StatusUnauthorized, map[string]string{"error": message})
}

func getUserID(sessionID string) (int32, error) {
	r, err := authClient.IsAuthenticated(context.Background(), &authpb.IsAuthenticatedRequest{SessionID: sessionID})
	if err != nil {
		app.errorLog.Printf("Cookie authentication check failed, error: %v", err)
		return 0, err
	}
	return r.UserID, nil
}
