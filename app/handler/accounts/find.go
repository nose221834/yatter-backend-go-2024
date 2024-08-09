package accounts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type FindRequest struct {
	Username string
}

func (h *handler) FindByUsername(w http.ResponseWriter, r *http.Request) {

	// URLパラメータからusernameを取得
	username := chi.URLParam(r, "username")
	if username == "" {
		http.Error(w, "username is required", http.StatusBadRequest)
		return
	}
	fmt.Println("Value of x:", username)

	ctx := r.Context()

	dto, err := h.accountUsecase.FindByUsername(ctx, username)
	if err != nil {
		if dto == nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dto.Account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
