package main

import (
	"convo_backend/handlers"
	"convo_backend/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)

	// middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{}))
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(handlers.Key),
		Skipper: func(c echo.Context) bool {
			if c.Path() == "/login" || c.Path() == "/signup" {
				return true
			}
			return false
		},
	}))

	// db connection
	db, err := mgo.Dial("localhost")

	if err != nil {
		e.Logger.Fatal(err)
	}

	// initialize unique indexes
	if err = db.Copy().DB("convo").C("users").EnsureIndex(mgo.Index{
		Key:    []string{"number"},
		Unique: true,
	}); err != nil {
		log.Fatal(err)
	}

	h := &handlers.Handler{DB: db}

	// routes setup
	routes.Setup(h, e)

	// start server
	e.Logger.Fatal(e.Start(":8000"))
}
