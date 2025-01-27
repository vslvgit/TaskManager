package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {

	var TaskInMap = map[int]Task{
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

	r.Get("/task/", GetAllTask)
	r.Get("/task/{id}", GetIDTask)
	r.Post("/task/", PostTask)
	r.Delete("/task/{id}", DeleteIDTask)
	r.Put("/task/{id}", PutIDTask)

	http.ListenAndServe(":8080", r)

}
