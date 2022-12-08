package router

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/ikalemmon/Gocha/handler"
	"github.com/ikalemmon/Gocha/service"
)

func NewRouter(db *sql.DB) *mux.Router {
	// register routes
	mux := mux.NewRouter()
	mux.HandleFunc("/healthz", handler.NewHealthzHandler().ServeHTTP)
	mux.HandleFunc("/result", handler.NewResultHandler(service.NewResultService(db)).ServeHTTP)
	//mux.HandleFunc("/user", handler.NewUserHandler(service.NewUserService(db)).ServeHTTP)
	return mux
}
