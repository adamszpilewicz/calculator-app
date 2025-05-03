package handler

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"

	"my-order-app/internal/calculator"
)

// Handler is the HTTP handler for managing pack sizes and calculating orders
type Handler struct {
	PackSizes []int
}

// NewHandler creates a new Handler with the given pack sizes
func NewHandler(packSizes []int) *Handler {
	return &Handler{PackSizes: packSizes}
}

// Calculate handles the calculation of packs based on the quantity provided in the query string
func (h *Handler) Calculate(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("quantity")
	quantity, err := strconv.Atoi(q)
	if err != nil || quantity <= 0 {
		http.Error(w, "Invalid quantity", http.StatusBadRequest)
		return
	}

	if len(h.PackSizes) == 0 {
		http.Error(w, "No pack sizes configured", http.StatusBadRequest)
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

// GetPackSizes returns the current pack sizes in JSON format
func (h *Handler) GetPackSizes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.PackSizes)
}

// AddPackSize adds a new pack size to the list, ensuring no duplicates
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

// DeletePackSize removes a pack size from the list
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

// DeleteAllPackSizes clears all pack sizes
func (h *Handler) DeleteAllPackSizes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	h.PackSizes = []int{} // clear pack sizes
	w.WriteHeader(http.StatusOK)
}
