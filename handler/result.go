package handler

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/ikalemmon/Gocha/model"
	"github.com/ikalemmon/Gocha/service"
)

type ResultHandler struct {
	svc *service.ResultService
}

func NewResultHandler(svc *service.ResultService) *ResultHandler {
	return &ResultHandler{
		svc: svc,
	}
}

func (h *ResultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/result.html"))

	request := &model.ResultRequest{}

	query := r.URL.Query()
	request.Size, _ = strconv.ParseInt(query.Get("size"), 10, 64)

	if request.Size == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := h.svc.GetResult(r.Context(), request.Size)
	if err != nil {
		log.Println(err)
		return
	}
	response := model.ResultResponse{Result: result}
	if err := t.ExecuteTemplate(w, "result.html", response); err != nil {
		log.Fatal(err)
		return
	}
}
