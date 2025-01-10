package receipts

import (
	"backend/utils"
	"encoding/json"
	"go/types"
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
	router.HandleFunc("/receipts/process", h.handleProcess).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", h.handlePoints).Methods("GET")
}

func (h *Handler) handleProcess(w http.ResponseWriter, r *http.Request) {
	var payload types.ReceiptPayload
	if err := utils.ParsePayload(r *http.Request, payload json){
		utils.WriteError(w,http.StatusBadRequest,err)
	}
	log.Println("Post working")
}

func (h *Handler) handlePoints(w http.ResponseWriter, r *http.Request) {
	log.Println("Get working")
}
