package symmdiff

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strings"
	"testing"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyz"
)

var (
	cases = []testCase{
		{
			name:           "both are empty",
			first:          []int{},
			second:         []int{},
			firstExpected:  []int{},
			secondExpected: []int{},
		},
		{
			name:           "second is empty",
			first:          []int{1, 2, 3},
			second:         []int{},
			firstExpected:  []int{1, 2, 3},
			secondExpected: []int{},
		},
		{
			name:           "first is empty",
			first:          []int{},
			second:         []int{1, 2, 3},
			firstExpected:  []int{},
			secondExpected: []int{1, 2, 3},
		},
		{
			name:           "both have same elements",
			first:          []int{1, 2, 3},
			second:         []int{1, 2, 3},
			firstExpected:  []int{},
			secondExpected: []int{},
		},
		{
			name:           "two elements are common",
			first:          []int{1, 2, 3},
			second:         []int{2, 3, 4},
			firstExpected:  []int{1},
			secondExpected: []int{4},
		},
		{
			name:           "no common elements",
			first:          []int{1, 2, 3},
			second:         []int{4, 5, 6},
			firstExpected:  []int{1, 2, 3},
			secondExpected: []int{4, 5, 6},
		},
		{
			name:           "first is subset of second",
			first:          []int{1, 2, 3},
			second:         []int{1, 2, 3, 4, 5, 6},
			firstExpected:  []int{},
			secondExpected: []int{4, 5, 6},
		},
		{
			name:           "second is subset of first",
			first:          []int{1, 2, 3, 4, 5, 6},
			second:         []int{1, 2, 3},
			firstExpected:  []int{4, 5, 6},
			secondExpected: []int{},
		},
		{
			name:           "both have duplicates and common elements",
			first:          []int{1, 1, 2},
			second:         []int{2, 3, 3},
			firstExpected:  []int{1},
			secondExpected: []int{3},
		},
	}
	r = rand.New(rand.NewSource(42))
	f = createStrings(100_000, 6)
	s = createStrings(100_000, 6)
)

func createStrings(amount, length int) []string {
	result := make([]string, amount)
	for i := 0; i < amount; i++ {
		result[i] = createString(r, length)
	}
	return result
}

func createString(r *rand.Rand, length int) string {
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteByte(charset[r.Intn(len(charset))])
	}
	return sb.String()
}

func createInts(amount, limit int) []int {
	result := make([]int, amount)
	for i := 0; i < amount; i++ {
		result[i] = r.Intn(limit)
	}
	return result
}

type testCase struct {
	name           string
	first          []int
	second         []int
	firstExpected  []int
	secondExpected []int
}

func TestBasicSymmDiff(t *testing.T) {
	for _, c := range cases {
		first, second := BasicSymmDiff(c.first, c.second)
		assert.ElementsMatch(t, c.firstExpected, first, "case: %s", c.name)
		assert.ElementsMatch(t, c.secondExpected, second, "case: %s", c.name)
	}
}

func TestBetterSymmDiff(t *testing.T) {
	for _, c := range cases {
		first, second := BetterSymmDiff(c.first, c.second)
		assert.ElementsMatch(t, c.firstExpected, first)
		assert.ElementsMatch(t, c.secondExpected, second)
	}
}

func TestSparseSymmDiff(t *testing.T) {
	for _, c := range cases {
		first, second := SparseSymmDiff(c.first, c.second)
		assert.ElementsMatch(t, c.firstExpected, first)
		assert.ElementsMatch(t, c.secondExpected, second)
	}
}

func BenchmarkBasicSymmDiff(b *testing.B) {
	for b.Loop() {
		BasicSymmDiff(f, s)
	}
}

func BenchmarkBetterSymmDiff(b *testing.B) {
	for b.Loop() {
		BetterSymmDiff(f, s)
	}
}

func BenchmarkSparseSymmDiff(b *testing.B) {
	for b.Loop() {
		SparseSymmDiff(f, s)
	}
}
