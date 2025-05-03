package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"my-order-app/internal/calculator"
)

func TestHandler_Calculate(t *testing.T) {
	h := NewHandler([]int{250, 500, 1000})

	req := httptest.NewRequest("GET", "/api/calculate?quantity=1201", nil)
	w := httptest.NewRecorder()

	h.Calculate(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	var body struct {
		Packs      []calculator.PackResult `json:"packs"`
		TotalItems int                     `json:"total_items"`
	}
	err := json.NewDecoder(resp.Body).Decode(&body)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, body.TotalItems, 1201)
	assert.NotEmpty(t, body.Packs)
}

func TestHandler_Calculate_InvalidQuantity(t *testing.T) {
	h := NewHandler([]int{250, 500, 1000})

	req := httptest.NewRequest("GET", "/api/calculate?quantity=abc", nil)
	w := httptest.NewRecorder()

	h.Calculate(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestHandler_GetPackSizes(t *testing.T) {
	initialSizes := []int{5000, 2000, 1000}
	h := NewHandler(initialSizes)

	req := httptest.NewRequest("GET", "/api/packsizes", nil)
	w := httptest.NewRecorder()

	h.GetPackSizes(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	var sizes []int
	err := json.NewDecoder(resp.Body).Decode(&sizes)
	assert.NoError(t, err)
	assert.Equal(t, initialSizes, sizes)
}

func TestHandler_AddPackSize(t *testing.T) {
	h := NewHandler([]int{1000})

	payload := map[string]int{"size": 2000}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest("POST", "/api/packsizes/add", bytes.NewReader(body))
	w := httptest.NewRecorder()

	h.AddPackSize(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Contains(t, h.PackSizes, 2000)
}

func TestHandler_AddPackSize_Duplicate(t *testing.T) {
	h := NewHandler([]int{1000})

	payload := map[string]int{"size": 1000}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest("POST", "/api/packsizes/add", bytes.NewReader(body))
	w := httptest.NewRecorder()

	h.AddPackSize(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestHandler_DeletePackSize(t *testing.T) {
	h := NewHandler([]int{1000, 2000})

	req := httptest.NewRequest("DELETE", "/api/packsizes/delete?size=1000", nil)
	w := httptest.NewRecorder()

	h.DeletePackSize(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotContains(t, h.PackSizes, 1000)
	assert.Contains(t, h.PackSizes, 2000)
}

func TestHandler_DeletePackSize_NotFound(t *testing.T) {
	h := NewHandler([]int{1000})

	req := httptest.NewRequest("DELETE", "/api/packsizes/delete?size=500", nil)
	w := httptest.NewRecorder()

	h.DeletePackSize(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestHandler_DeleteAllPackSizes(t *testing.T) {
	h := NewHandler([]int{500, 1000, 2000})

	req := httptest.NewRequest("DELETE", "/api/packsizes/deleteall", nil)
	w := httptest.NewRecorder()

	h.DeleteAllPackSizes(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Empty(t, h.PackSizes)
}
