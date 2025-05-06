package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"hrorscope/internal/renderers"
	"hrorscope/pkg/routes"
	"log"
)

func main() {

	err1 := godotenv.Load("api.env")
	if err1 != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))

	e.Renderer = &renderers.HTMLTemplate{
		Dir: "templates",
		Ext: ".html",
	}

	e.Static("/css", "css")
	e.Static("/img", "templates/img")
	routes.SetupRoutes(e)

	err := e.Start(":8080")

	if err != nil {
		log.Fatal("the server isnt running: %v", err)
	}

}
