package symmdiff

func BasicSymmDiff[T comparable](first, second []T) ([]T, []T) {
	return diff(first, second), diff(second, first)
}

func diff[T comparable](first, second []T) []T {
	seen := make(map[T]bool, len(first)+len(second))
	for _, e := range second {
		seen[e] = true
	}
	result := make([]T, 0, len(first))
	for _, item := range first {
		if !seen[item] {
			result = append(result, item)
			seen[item] = true
		}
	}
	return result
}

func BetterSymmDiff[T comparable](first, second []T) ([]T, []T) {
	seen := make(map[T]bool, len(first)+len(second))
	for _, e := range first {
		seen[e] = false
	}
	secondUnique := make([]T, 0, len(second))
	for _, e := range second {
		if _, ok := seen[e]; ok {
			seen[e] = true
		} else {
			seen[e] = false
			secondUnique = append(secondUnique, e)
		}
	}
	firstUnique := make([]T, 0, len(first))
	for _, e := range first {
		if !seen[e] {
			seen[e] = true
			firstUnique = append(firstUnique, e)
		}
	}
	return firstUnique, secondUnique
}

func SparseSymmDiff[T comparable](first, second []T) ([]T, []T) {
	firstUnique := newSparseSet[T](len(first))
	for _, e := range first {
		firstUnique.add(e)
	}
	seen := make(map[T]bool, len(second))
	secondUnique := make([]T, 0, len(second))
	for _, e := range second {
		if seen[e] {
			continue
		}
		if firstUnique.has(e) {
			firstUnique.remove(e)
		} else {
			secondUnique = append(secondUnique, e)
		}
		seen[e] = true
	}
	return firstUnique.dense, secondUnique
}
