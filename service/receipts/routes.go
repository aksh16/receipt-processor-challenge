package receipts

import (
	"backend/types"
	"backend/utils"
	"log"
	"net/http"
	"strconv"

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
		log.Println("Incorrect format")
		utils.WriteResponse(w, http.StatusBadRequest, "The receipt is invalid")
		return
	}
	emptyFields := utils.GetEmptyJSONFields(payload)
	if len(emptyFields) > 0 {
		log.Println("Empty fields")
		utils.WriteResponse(w, http.StatusBadRequest, "The receipt is invalid")
		return
	}
	points := utils.CalculatePoints(payload)
	receipt_id, err := h.store.AddPoints(points)
	if err != nil {
		log.Println("Failed to add to db")
		utils.WriteResponse(w, http.StatusBadRequest, "The receipt is invalid")
		return
	}
	id_response := types.IdResponse{Id: receipt_id}
	utils.WriteResponse(w, http.StatusOK, id_response)
}

func (h *Handler) handlePoints(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)
	receipt_id, err := strconv.ParseUint(id["id"], 10, 64)
	if err != nil {
		log.Println("Invalid ID format")
		utils.WriteResponse(w, http.StatusBadRequest, "No receipt found for that id")
		return
	}
	points, err := h.store.GetPoints(receipt_id)
	if err != nil {
		log.Printf("ID:%d, No record in db", receipt_id)
		utils.WriteResponse(w, http.StatusBadRequest, "No receipt found for that id")
		return
	}
	points_response := types.PointsResponse{Points: points}
	utils.WriteResponse(w, http.StatusOK, points_response)
}
