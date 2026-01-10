package main

import (
	"database/sql"
	"embed"
	"log/slog"
	"meal-plan/internal/config"
	"meal-plan/internal/database"
	"meal-plan/internal/handlers"
	"meal-plan/internal/repository"
	"meal-plan/internal/service"
	"net/http"
	"os"
)

var migrationsFS embed.FS

func main() {
	cfg := config.Load()
	setupLogger()

	slog.Info("starting server meal plan", "port", cfg.Port)

	db, err := database.New(cfg.GetDNS())
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	application := mapHandlers(db)
	runServer(cfg.Port, application)
}

func mapHandlers(db *sql.DB) *handlers.MealPlanHandler {
	repo := repository.NewMealPlanRepository(db)
	srv := service.NewMealPlanService(repo)
	return handlers.NewMealPlanHandler(srv)
}

func setupLogger() {
	opts := &slog.HandlerOptions{Level: slog.LevelDebug}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	slog.SetDefault(logger)
}

func runServer(port string, h *handlers.MealPlanHandler) {
	slog.Info("server started", "address", "http://localhost:"+port)
	if err := http.ListenAndServe(":"+port, h.InitRoutes()); err != nil {
		slog.Error("server failed", "error", err)
		os.Exit(1)
	}
}
