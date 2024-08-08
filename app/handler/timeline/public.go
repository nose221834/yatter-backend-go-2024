package timeline

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *handler) Public(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")

	// デフォルト値の設定
	limit := 10

	// クエリパラメータを変換
	if limitStr != "" {
		var err error
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
			return
		}
	}

	ctx := r.Context()

	dto, err := h.tu.Public(ctx, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dto.Timeline); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
