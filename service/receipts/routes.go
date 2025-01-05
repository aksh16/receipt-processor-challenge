package receipts

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/receipts/process", h.handleProcess).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", h.handlePoints).Methods("GET")
}

func (h *Handler) handleProcess(w http.ResponseWriter, r *http.Request) {
	log.Println("Post working")
}

func (h *Handler) handlePoints(w http.ResponseWriter, r *http.Request) {
	log.Println("Get working")
}
