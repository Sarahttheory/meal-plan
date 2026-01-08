package main

import (
	"log"
	"meal-plan/internal/config"
	"meal-plan/internal/handlers"
	"meal-plan/internal/repository"
	"meal-plan/internal/service"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.Load()
	db := repository.InitDB()
	defer db.Close()

	repo := &repository.MealPlanRepository{DB: db}
	srv := &service.MealPlanService{Repo: repo}
	h := &handlers.MealPlanHandler{Service: srv}

	router := h.InitRoutes()

	log.Printf("Starting server on port %s", cfg.Port)
	//if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
	//	log.Fatal(err)
	//} //TODO временно, пока порт занят
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
