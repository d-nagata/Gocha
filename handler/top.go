package handler

import (
	"html/template"
	"net/http"
)

type TopHandler struct {
}

func NewTopHandler() *TopHandler {
	return &TopHandler{}
}

func (h *TopHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/top.html"))
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}
