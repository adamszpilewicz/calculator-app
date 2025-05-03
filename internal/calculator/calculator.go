package calculator

type PackResult struct {
	Size  int `json:"size"`
	Count int `json:"count"`
}

func CalculatePacks(quantity int, packSizes []int) []PackResult {
	best := []PackResult{}
	bestItems := 0
	bestCount := 0

	var helper func(index int, current []PackResult, total int)
	helper = func(index int, current []PackResult, total int) {
		if total >= quantity {
			if bestItems == 0 || total < bestItems || (total == bestItems && totalPacks(current) < bestCount) {
				bestItems = total
				best = append([]PackResult(nil), current...)
				bestCount = totalPacks(current)
			}
			return
		}
		if index >= len(packSizes) {
			return
		}

		size := packSizes[index]
		maxCount := (quantity + size - 1) / size

		for count := 0; count <= maxCount; count++ {
			next := append([]PackResult(nil), current...)
			if count > 0 {
				next = append(next, PackResult{Size: size, Count: count})
			}
			helper(index+1, next, total+count*size)
		}
	}
	helper(0, []PackResult{}, 0)
	return best
}

func totalPacks(packs []PackResult) int {
	total := 0
	for _, p := range packs {
		total += p.Count
	}
	return total
}

func TotalItems(packs []PackResult) int {
	total := 0
	for _, p := range packs {
		total += p.Size * p.Count
	}
	return total
}
