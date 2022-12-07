package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/ikalemmon/Gocha/model"
)

type HealthzHandler struct{}

func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{}
}

func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/healthz.html"))

	var response = &model.HealthzResponse{}
	(*response).Message = "OK"

	if err := t.ExecuteTemplate(w, "healthz.html", *response); err != nil {
		log.Fatal(err)
	}

	// err := json.NewEncoder(w).Encode(*response)
	// if err != nil {
	// 	log.Println(err)
	// }
}
