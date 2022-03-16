package tronics

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func middlewareSample(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("inside the middle ware")
		return next(c)
	}
}
