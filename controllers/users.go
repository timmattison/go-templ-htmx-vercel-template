package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/timmattison/go-templ-htmx-vercel-template/views/users"
)

type UserHandler struct{}

func (h UserHandler) Index(c echo.Context) error {
	return render(c, users.Index())
}
