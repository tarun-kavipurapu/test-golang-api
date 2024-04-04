package main

import (
	"log"

	"github.com/tarun-kavipurapu/gin-gorm/db"
	"github.com/tarun-kavipurapu/gin-gorm/db/models"
	"github.com/tarun-kavipurapu/gin-gorm/routes"
)

func loadDatabase() {
	db.InitDb()
	err := db.DB.AutoMigrate(&models.User{}, &models.Role{}, &models.Token{}, &models.Address{}, &models.Restaurant{})
	if err != nil {
		panic("[Error] inn migration " + err.Error())
	}

}
func serveApplication() {
	r := routes.SetupRouter()
	err := r.Run(":8000")
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
	}
}

func main() {

	// loadEnv()
	loadDatabase()
	log.Print("I am here in main")
	serveApplication()

}

/*1. setup Routes
2. Setup Repositories
3. Start the authentication with the JWTS
*/
