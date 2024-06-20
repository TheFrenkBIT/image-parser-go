package main

import (
	"github.com/labstack/echo"
	"image-parser/app/handlers"
)

func main() {
	e := echo.New()
	e.POST("/parse", handlers.ImageHandler())
	e.Start(":80")
}
