package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	size := 100
	result := generateRandomElements(size)
	if len(result) != size {
		t.Errorf("ожидалась длина слайса %d, получена %d", size, len(result))
	}

	emptyResult := generateRandomElements(0)
	if len(emptyResult) != 0 {
		t.Errorf("ожидался пустой слайс при size=0, получена длина %d", len(emptyResult))
	}

	negativeResult := generateRandomElements(-5)
	if len(negativeResult) != 0 {
		t.Errorf("ожидался пустой слайс при отрицательном size, получена длина %d", len(negativeResult))
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "пустой слайс",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "nil слайс",
			input:    nil,
			expected: 0,
		},
		{
			name:     "один элемент",
			input:    []int{42},
			expected: 42,
		},
		{
			name:     "несколько элементов",
			input:    []int{10, 55, 23, 4},
			expected: 55,
		},
		{
			name:     "одинаковые элементы",
			input:    []int{7, 7, 7, 7},
			expected: 7,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := maximum(tc.input)
			if got != tc.expected {
				t.Errorf("для %s ожидалось %d, получено %d", tc.name, tc.expected, got)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "пустой слайс",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "nil слайс",
			input:    nil,
			expected: 0,
		},
		{
			name:     "один элемент",
			input:    []int{42},
			expected: 42,
		},
		{
			name:     "меньше элементов, чем CHUNKS",
			input:    []int{10, 99, 5},
			expected: 99,
		},
		{
			name:     "кратно CHUNKS",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			expected: 16,
		},
		{
			name:     "некратно CHUNKS, максимум в самом хвосте",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 500},
			expected: 500,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := maxChunks(tc.input)
			if got != tc.expected {
				t.Errorf("для %s ожидалось %d, получено %d", tc.name, tc.expected, got)
			}
		})
	}

	t.Run("сравнение maximum и maxChunks", func(t *testing.T) {
		data := generateRandomElements(10_000)
		expected := maximum(data)
		got := maxChunks(data)
		if got != expected {
			t.Errorf("результаты не совпадают: maximum=%d, maxChunks=%d", expected, got)
		}
	})
}
