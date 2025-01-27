package main

import (
	"net/http"
	"time"

	h "TaskManager/handlers"
	"TaskManager/stracts"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {

	var Tasks = map[int]stracts.Task{
		1: {
			ID:        1,
			Title:     "Make me happy!",
			Completed: true,
			CreatedAt: time.Now(),
		},
		2: {
			ID:        2,
			Title:     "Make sure I am American!",
			Completed: false,
			CreatedAt: time.Now(),
		},
		3: {
			ID:        3,
			Title:     "Make my dream come true!",
			Completed: true,
			CreatedAt: time.Now(),
		},
		4: {
			ID:        4,
			Title:     "I want to live in peace!",
			Completed: false,
			CreatedAt: time.Now(),
		},
	}
	r := chi.NewRouter()

	r.Use(middleware.Logger)    // Логирование запросов
	r.Use(middleware.Recoverer) // Восстановление после паник

	r.Get("/task/", h.GetAllTask)
	r.Get("/task/{id}", h.GetIDTask)
	r.Post("/task/", h.PostTask)
	r.Delete("/task/{id}", DeleteIDTask)
	r.Put("/task/{id}", PutIDTask)

	err = http.ListenAndServe(":8080", r)

}
