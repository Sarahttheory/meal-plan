package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (h *MealPlanHandler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/mealPlan", func(r chi.Router) {
		r.Get("/dishes/all", h.GetAllDishes)
	})

	return r
}
