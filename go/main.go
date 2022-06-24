package main

import (
	"net/http"
	"os"
	"strings"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		htmlMessage := "Hello, Docker!"

		fmt.Println()
    	for _, e := range os.Environ() {
	        pair := strings.SplitN(e, "=", 2)
			// 	# fmt.Println(pair[0])
			htmlMessage = htmlMessage + "<p>" + pair[0] + " :: " + pair[1] + " :: " + e
    	}
		// return c.HTML(http.StatusOK, "Hello, Docker! <3")
		return c.HTML(http.StatusOK, htmlMessage)
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "80"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}