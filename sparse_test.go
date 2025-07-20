package symmdiff

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSparseSet(t *testing.T) {
	t.Run("supports simple add, has, remove", func(t *testing.T) {
		set := newSparseSet[int](4)
		set.add(10)
		set.add(20)
		require.True(t, set.has(10))
		require.True(t, set.has(20))

		set.remove(10)
		require.False(t, set.has(10))
		require.True(t, set.has(20))

		set.remove(99)
		require.True(t, set.has(20))
	})

	t.Run("can't add duplicates", func(t *testing.T) {
		set := newSparseSet[int](4)
		set.add(1)
		set.add(1)
		require.Len(t, set.dense, 1)
	})

	t.Run("removes last element", func(t *testing.T) {
		set := newSparseSet[int](4)
		set.add(1)
		set.add(2)
		set.add(3)
		set.remove(3)
		require.Len(t, set.dense, 2)
		require.False(t, set.has(3))
		require.True(t, set.has(1))
		require.True(t, set.has(2))
	})

	t.Run("removes in different order", func(t *testing.T) {
		set := newSparseSet[int](4)
		set.add(1)
		set.add(2)
		set.add(3)
		set.remove(2)
		require.Len(t, set.dense, 2)
		require.False(t, set.has(2))
		require.True(t, set.has(1))
		require.True(t, set.has(3))

		set.remove(1)
		require.Len(t, set.dense, 1)
		require.False(t, set.has(1))
		require.True(t, set.has(3))

		set.remove(3)
		require.Len(t, set.dense, 0)
		require.False(t, set.has(3))
	})

	t.Run("adds again after removing", func(t *testing.T) {
		set := newSparseSet[int](3)
		set.add(1)
		set.remove(1)
		require.False(t, set.has(1))
		require.Len(t, set.dense, 0)
		set.add(1)
		require.True(t, set.has(1))
	})

	t.Run("contains remained elements after removing", func(t *testing.T) {
		set := newSparseSet[int](4)
		set.add(1)
		set.add(2)
		set.add(3)
		set.remove(2)
		set.add(4)
		require.ElementsMatch(t, []int{1, 3, 4}, set.dense)
	})
}
