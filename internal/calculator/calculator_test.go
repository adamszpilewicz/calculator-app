package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePacks(t *testing.T) {
	tests := []struct {
		name           string
		quantity       int
		packSizes      []int
		expectedResult []PackResult
		expectedTotal  int
	}{
		{
			name:      "Order exactly one pack",
			quantity:  250,
			packSizes: []int{250, 500, 1000},
			expectedResult: []PackResult{
				{Size: 250, Count: 1},
			},
			expectedTotal: 250,
		},
		{
			name:      "Order slightly above pack size",
			quantity:  251,
			packSizes: []int{250, 500, 1000},
			expectedResult: []PackResult{
				{Size: 500, Count: 1},
			},
			expectedTotal: 500,
		},
		{
			name:      "Order with multiple pack sizes",
			quantity:  501,
			packSizes: []int{250, 500, 1000},
			expectedResult: []PackResult{
				{Size: 500, Count: 1},
				{Size: 250, Count: 1},
			},
			expectedTotal: 750,
		},
		{
			name:      "Edge case large quantity",
			quantity:  500000,
			packSizes: []int{23, 31, 53},
			expectedResult: []PackResult{
				{Size: 53, Count: 9429},
				{Size: 31, Count: 7},
				{Size: 23, Count: 2},
			},
			expectedTotal: 500000 + 0, // actual total items will match or exceed quantity
		},
		{
			name:      "Order requiring all pack sizes",
			quantity:  12001,
			packSizes: []int{250, 500, 1000, 2000, 5000},
			expectedResult: []PackResult{
				{Size: 5000, Count: 2},
				{Size: 2000, Count: 1},
				{Size: 250, Count: 1},
			},
			expectedTotal: 12250,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculatePacks(tt.quantity, tt.packSizes)

			// Check total items
			totalItems := TotalItems(result)
			assert.GreaterOrEqual(t, totalItems, tt.quantity, "total items should fulfill quantity")
			assert.Equal(t, tt.expectedTotal, totalItems, "total items mismatch")

			// Check pack counts per size
			actualMap := make(map[int]int)
			for _, r := range result {
				actualMap[r.Size] = r.Count
			}
			expectedMap := make(map[int]int)
			for _, r := range tt.expectedResult {
				expectedMap[r.Size] = r.Count
			}

			assert.Equal(t, expectedMap, actualMap, "pack combinations mismatch")
		})
	}
}
