package receipts

import (
	"backend/types"
	"backend/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ReceiptStore
}

func NewHandler(store types.ReceiptStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/receipts/process", h.handleReceipts).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", h.handlePoints).Methods("GET")
}

func (h *Handler) handleReceipts(w http.ResponseWriter, r *http.Request) {
	var payload types.ReceiptPayload
	if err := utils.ParsePayload(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	log.Println("Post working")
}

func (h *Handler) handlePoints(w http.ResponseWriter, r *http.Request) {
	log.Println("Get working")
}
