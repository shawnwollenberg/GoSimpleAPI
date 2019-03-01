package shawn

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Todo struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{todoID}", GetATodo)
	router.Get("/", GetShawn)
	return router
}

func GetATodo(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	todos := Todo{
		Slug:  todoID,
		Title: "Hello World",
		Body:  "Hello World from planet Earth",
	}
	render.JSON(w, r, todos)
}
func GetShawn(w http.ResponseWriter, r *http.Request) {
	todos := Todo{
		Slug:  "Hey Shawn",
		Title: "Good Test",
		Body:  "Look at you go shawn",
	}
	render.JSON(w, r, todos)
}
