package symmdiff

type sparseSet[T comparable] struct {
	sparse map[T]int
	dense  []T
}

func newSparseSet[T comparable](size int) *sparseSet[T] {
	return &sparseSet[T]{
		sparse: make(map[T]int, size),
		dense:  make([]T, 0, size),
	}
}

func (s *sparseSet[T]) has(item T) bool {
	_, ok := s.sparse[item]
	return ok
}

func (s *sparseSet[T]) add(item T) {
	if !s.has(item) {
		s.sparse[item] = len(s.dense)
		s.dense = append(s.dense, item)
	}
}

func (s *sparseSet[T]) remove(item T) {
	if index, ok := s.sparse[item]; ok {
		last := s.dense[len(s.dense)-1]
		s.sparse[last] = index
		s.dense[index] = last
		delete(s.sparse, item)
		s.dense = s.dense[:len(s.dense)-1]
	}
}
