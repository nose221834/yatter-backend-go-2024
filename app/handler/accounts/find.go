package accounts

import (
	"encoding/json"
	"net/http"
)

type FindRequest struct {
	Username string
}

func (h *handler) Find(w http.ResponseWriter, r *http.Request) {
	var req FindRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	dto ,err := h.accountUsecase.Find(ctx,req.Username)
	if err := nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","application/json")
	if err := json.NewEncoder(w).Encode(dto.Account); err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}






}
