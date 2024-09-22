package main

import (
	"flag"
	"hwc/db"
	middlewares "hwc/middlewares"
	routes "hwc/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	seed := flag.Bool("seed", false, "seed the database")
	flag.Parse()

	if *seed {
		db.Db.Seed()
		return
	}

	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://nlw-pocket-web-production.up.railway.app", "http://localhost:5173"},
		AllowHeaders:     []string{"Content-Type"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	authenticated := router.Group("/")

	authenticated.Use(middlewares.AuthMiddleware())
	{
		authenticated.GET("/cars", routes.GetAllCars())
		authenticated.GET("/search", routes.SearchCars())
	}

	router.Run(":8000")
}
