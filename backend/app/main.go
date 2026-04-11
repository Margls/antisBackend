package main

import (
	"log"
	"net/http"

	"antis/backend/config"
	"antis/backend/database"
	"antis/backend/handlers"
	"antis/backend/repositories"
	"antis/backend/routes"
	"antis/backend/service"
)

func main() {
	config.LoadEnv()

	connStr, err := config.PostgresDSN()
	if err != nil {
		log.Fatalf("database configuration: %v", err)
	}

	db, err := database.NewDB(connStr)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	defer db.Close()

	
	log.Println("Server started on :8080")


	itemsRepo := repositories.NewItemRepository(db)
	itemService := service.NewItemService(itemsRepo)
	itemHandler := handlers.NewItemHandler(itemService)
	router := routes.SetupRouter(itemHandler)



    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatal(err)
    }
}