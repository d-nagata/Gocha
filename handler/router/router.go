package router

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/ikalemmon/Gocha/handler"
)

func NewRouter(db *sql.DB) *mux.Router {
	// register routes
	mux := mux.NewRouter()
	mux.HandleFunc("/healthz", handler.NewHealthzHandler().ServeHTTP)
	//mux.HandleFunc("/user", handler.NewUserHandler(service.NewUserService(db)).ServeHTTP)
	return mux
}
