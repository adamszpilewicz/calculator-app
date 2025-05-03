package handler

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"

	"my-order-app/internal/calculator"
)

type Handler struct {
	PackSizes []int
}

func NewHandler(packSizes []int) *Handler {
	return &Handler{PackSizes: packSizes}
}

func (h *Handler) Calculate(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("quantity")
	quantity, err := strconv.Atoi(q)
	if err != nil || quantity <= 0 {
		http.Error(w, "Invalid quantity", http.StatusBadRequest)
		return
	}

	packs := calculator.CalculatePacks(quantity, h.PackSizes)
	total := calculator.TotalItems(packs)

	resp := struct {
		Packs      []calculator.PackResult `json:"packs"`
		TotalItems int                     `json:"total_items"`
	}{
		Packs:      packs,
		TotalItems: total,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) GetPackSizes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.PackSizes)
}

func (h *Handler) AddPackSize(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload struct {
		Size int `json:"size"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil || payload.Size <= 0 {
		http.Error(w, "Invalid size", http.StatusBadRequest)
		return
	}

	// Check for duplicates
	for _, s := range h.PackSizes {
		if s == payload.Size {
			http.Error(w, "Pack size already exists", http.StatusBadRequest)
			return
		}
	}

	h.PackSizes = append(h.PackSizes, payload.Size)
	sort.Sort(sort.Reverse(sort.IntSlice(h.PackSizes)))

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeletePackSize(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	q := r.URL.Query().Get("size")
	size, err := strconv.Atoi(q)
	if err != nil || size <= 0 {
		http.Error(w, "Invalid size", http.StatusBadRequest)
		return
	}

	found := false
	newSizes := []int{}
	for _, s := range h.PackSizes {
		if s == size {
			found = true
		} else {
			newSizes = append(newSizes, s)
		}
	}

	if !found {
		http.Error(w, "Pack size not found", http.StatusNotFound)
		return
	}

	h.PackSizes = newSizes
	w.WriteHeader(http.StatusOK)
}
