package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	authpb "github.com/Caps1d/task-manager-cloud-app/auth/pb"
	userpb "github.com/Caps1d/task-manager-cloud-app/user/pb"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Handlers
func home(c echo.Context) error {
	// 1. If user is not logged in -> show hero landing page or whatever its called, include a login form with register option
	// -> Handle auth with gRPC requests to Auth service
	// 2.0 If user is logged in -> show team space

	return c.String(http.StatusOK, "Home Page!")
}

func getLogin(c echo.Context) error {
	// check if user-session cookie exists
	cookie, err := c.Cookie("user-session")
	if err != nil && !errors.Is(err, http.ErrNoCookie) {
		app.infoLog.Printf("%v", err)
	}
	// check if sessionID matches
	if cookie != nil {
		userID, err := getUserID(cookie.Value)
		if err != nil {
			app.errorLog.Printf("API: error at getLogin when fetching userID %v", err)
			return err
		}
		if userID != 0 {
			app.infoLog.Printf("User logged in, redirecting to home page")
			c.Redirect(http.StatusPermanentRedirect, "/")
			return nil
		}
	}

	return c.String(http.StatusOK, "User endpoint reached")
}

type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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

type RegistrationForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	Name     string `json:"name"`
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
	Id       int32  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Role     string `json:"role"`
	TeamId   int32  `json:"teamId"`
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

type TeamForm struct {
	Name        string `json:"name"`
	ManagerName string `json:"managerName"`
}

func postTeam(c echo.Context) error {
	var form TeamForm

	err := c.Bind(&form)
	if err != nil {
		app.errorLog.Printf("API: failed to bind json data at postTeam %v", err)
		return app.badRequest(c, "Invalid request data")
	}

	cookie, err := c.Cookie("user-session")
	if err != nil {
		app.errorLog.Printf("API: failed to read user-session cookie at postTeam %v", err)
		return app.unauthorized(c, "Unauthorized request")
	}

	userID, err := getUserID(cookie.Value)
	if err != nil {
		app.errorLog.Printf("API: failed to get userID from session at postTeam %v", err)
		return app.serverError(c, "Internal server error")
	}

	r, err := userClient.CreateTeam(context.Background(), &userpb.CreateTeamRequest{Name: form.Name, Manager: userID})
	if err != nil {
		app.errorLog.Printf("API: failed at CreateTeam pb request at postTeam %v", err)
		return app.serverError(c, "Failed to create team")
	}

	app.infoLog.Printf("API: Success team %v created!", form.Name)

	// review links
	data := map[string]interface{}{
		"message": "Team created successfully!",
		"team": map[string]interface{}{
			"teamID":      r.Id,
			"teamName":    form.Name,
			"managerName": form.ManagerName,
		},
		"links": map[string]interface{}{
			"self":    fmt.Sprintf("/teams/%d", r.Id),
			"members": fmt.Sprintf("/teams/%d/members", r.Id),
		},
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return c.JSON(http.StatusCreated, data)
}

type Member struct {
	ID       int32
	Name     string
	Email    string
	Username string
	Role     string
}

type Team struct {
	ID      int32    `json:"id"`
	Name    string   `json:"name"`
	Manager int32    `json:"manager"`
	Members []Member `json:"members"`
	Size    int32    `json:"size"`
}

func getTeam(c echo.Context) error {
	if allowed, err := isAuthorized(c); !allowed {
		app.errorLog.Printf("API: unauthorized route access at getTeam %v", err)
		return app.unauthorized(c, "Unauthorized request")
	}

	// sanitize query params
	val := c.QueryParam("id")
	teamID, _ := strconv.Atoi(val)

	r, err := userClient.GetTeam(context.Background(), &userpb.GetTeamRequest{Id: int32(teamID)})
	if err != nil {
		app.errorLog.Printf("API: failed at GetTeam pb request %v", err)
		return app.badRequest(c, "Invalid request data")
	}

	team := r.GetTeam()
	members := team.GetMembers()

	t := &Team{
		ID:      team.Id,
		Name:    team.Name,
		Manager: team.Manager,
	}

	for _, member := range members {
		t.Members = append(t.Members, Member{ID: member.Id, Name: member.Name, Email: member.Email, Username: member.Username, Role: member.Role})
	}

	t.Size = int32(len(t.Members))

	resp := &Response{
		Message: fmt.Sprintf("Got team %v", teamID),
		Data:    t,
	}

	return c.JSON(http.StatusOK, resp)
}

func putTeam(c echo.Context) error {
	var team Team

	if allowed, err := isAuthorized(c); !allowed {
		app.errorLog.Printf("API: unauthorized route access at getTeam %v", err)
		return app.unauthorized(c, "Unauthorized request")
	}

	err := c.Bind(&team)
	if err != nil {
		app.errorLog.Printf("API: failed to bind request data at putTeam %v", err)
		return app.serverError(c, "Internal server error")
	}

	_, err = userClient.UpdateTeam(context.Background(), &userpb.UpdateTeamRequest{Id: team.ID, Name: &team.Name, Manager: &team.Manager, UserId: &team.Members[0].ID, Role: &team.Members[0].Role})
	if err != nil {
		app.errorLog.Printf("API: failed at UpdateTeam pb request %v", err)
		return app.badRequest(c, "Invalid request data")
	}

	return c.String(http.StatusOK, fmt.Sprintf("Team %v successfully updated", team.ID))
}
