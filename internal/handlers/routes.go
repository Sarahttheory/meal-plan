package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (h *MealPlanHandler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/index.html")
	})
	r.Get("/dashboard", h.GetDashboard)
	r.Route("/plan", func(r chi.Router) {
		r.Get("/weekly", h.GetWeeklyPlan)
		r.Put("/item", h.SaveItem)
	})

	r.Route("/dishes", func(r chi.Router) {
		r.Get("/", h.GetDishes)
		r.Put("/", h.SaveDish)
	})
	r.Route("/ingredients", func(r chi.Router) {
		r.Get("/", h.GetIngredients)
		r.Put("/", h.SaveIngredient)
	})

	return r
}
