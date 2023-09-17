package app

import (
	"net/http"

	"github.com/an33k/go-microservice/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func loadRoutes () *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	
	router.Get("/", func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/orders", loadOrderRoutes)

	return router
}

func loadOrderRoutes(router chi.Router) {
	orderHandlers := &handler.Order{}

	router.Post("/", orderHandlers.CreateOrder)
	router.Get("/", orderHandlers.ListOrder)
	router.Get("/{id}", orderHandlers.GetByID)
	router.Put("/{id}", orderHandlers.UpdateByID)
	router.Delete("/{id}", orderHandlers.DeleteByID)
}