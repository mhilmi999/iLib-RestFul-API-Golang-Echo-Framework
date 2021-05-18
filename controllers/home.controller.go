package controllers

import (
	. "mhilmi999/project-2-mhilmi999/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomeView(c echo.Context) error {
	flash := c.(*CustomContext).GetFlash()
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"flash": flash,
	})
}

func LoginView(c echo.Context) error {
	flash := c.(*CustomContext).GetFlash()
	return c.Render(http.StatusOK, "login.html", map[string]interface{}{
		"flash": flash,
	})
}
