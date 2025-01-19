package main

import (
	"github.com/labstack/echo/v4"

	"github.com/Yamako76/go-react/domain"
	"github.com/Yamako76/go-react/usecase"
)

type Registry struct {
	UserRepository domain.UserRepository
}

func RegisterRoutes(e *echo.Echo, registry Registry) {
	e.GET("/users/:id", func(c echo.Context) error {
		return usecase.GetUser(c, registry.UserRepository)
	})
}
