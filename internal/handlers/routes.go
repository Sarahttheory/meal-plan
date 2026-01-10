package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (h *MealPlanHandler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/index.html")
	})
	r.Route("/plan", func(r chi.Router) {
		r.Get("/weekly", h.GetWeeklyPlan)
		//r.Put("/add")
	})

	r.Route("/dishes", func(r chi.Router) {
		r.Get("/", h.GetDishes)
		r.Put("/", h.SaveDish)
		r.Get("/ingredients", h.GetIngredients)
	})

	return r
}
