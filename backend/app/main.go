package main

import (
	"log"
	"net/http"
	"os"
	"antis/backend/service"
	"antis/backend/repositories"
	"antis/backend/routes"
	"github.com/joho/godotenv"
	"antis/backend/handlers"
)

func main() {


	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST") 
    port := os.Getenv("DB_PORT")
    dbname := os.Getenv("DB_NAME")
    sslmode := os.Getenv("DB_SSLMODE")

	
	log.Println("Server started on :8080")


	itemsRepo := repositories.NewItemRepository()
	itemService := service.NewItemService()
	itemHandler := handlers.NewItemHandler()
	router := routes.SetupRouter()



    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatal(err)
    }
}