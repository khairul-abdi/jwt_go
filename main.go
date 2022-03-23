package main

import (
	"myGram/database"
	"myGram/router"
)

func main() {
	database.StartDB()

	r := router.StartApp()
	r.Run(":8080")
}
