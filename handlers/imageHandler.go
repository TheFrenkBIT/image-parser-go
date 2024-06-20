package handlers

import (
	"github.com/labstack/echo"
	"image-parser/app/repositories/core"
	"image-parser/app/services"
	"net/http"
)

func ImageHandler() echo.HandlerFunc {
	service := services.ParseService{
		Db: core.GetConnect(),
	}
	return func(context echo.Context) error {
		return context.JSON(http.StatusOK, service.Parse(context))
	}
}
