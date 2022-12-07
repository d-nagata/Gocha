package router

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	// register routes
	mux := mux.NewRouter()

	//mux.HandleFunc("/user", handler.NewUserHandler(service.NewUserService(db)).ServeHTTP)
	return mux
}
