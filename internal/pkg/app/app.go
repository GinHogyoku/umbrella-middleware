package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"umbrella-middleware/internal/app/endpoint"
	"umbrella-middleware/internal/app/mw"
	"umbrella-middleware/internal/app/service"
)

type App struct {
	e *endpoint.Endpoint
	s *service.Service
	r *mux.Router
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()

	a.e = endpoint.New(a.s)

	a.r = mux.NewRouter()

	a.r.Use(mw.RoleCheck)
	a.r.HandleFunc("/status", a.e.Status).Methods("GET")

	return a, nil

}

func (a *App) Run() error {
	log.Println("starting server on :8080")
	if err := http.ListenAndServe(":8080", a.r); err != nil {
		log.Fatal(err)
	}
	return nil
}
