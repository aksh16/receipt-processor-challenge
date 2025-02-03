package receipts

import (
	"backend/types"
	"backend/utils"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

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
	emptyFields := utils.GetEmptyJSONFields(payload)
	if len(emptyFields) > 0 {
		utils.WriteError(w, http.StatusBadRequest, errors.New(fmt.Sprintf("empty fields found: %s",
			strings.Join(emptyFields, ", "))))
	}
	points := utils.CalculatePoints(payload)

	points_response := map[string]uint64{"points": points}
	utils.WriteResponse(w, http.StatusOK, points_response)
}

func (h *Handler) handlePoints(w http.ResponseWriter, r *http.Request) {
	log.Println("Get working")
}
