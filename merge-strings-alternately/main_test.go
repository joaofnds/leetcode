package main

import (
	"fmt"
	"testing"
)

func TestMergeAlternately(t *testing.T) {
	tests := []struct {
		a        string
		b        string
		expected string
	}{
		{"abc", "pqr", "apbqcr"},
		{"ab", "pqrs", "apbqrs"},
		{"abcd", "pq", "apbqcd"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("mergeAlternately(%q,%q)", test.a, test.b), func(t *testing.T) {
			actual := mergeAlternately(test.a, test.b)
			if actual != test.expected {
				t.Errorf("expected %q, got %q", test.expected, actual)
			}
		})
	}
}

func BenchmarkMergeAlternately(b *testing.B) {
	for b.Loop() {
		mergeAlternately("abcde", "12345")
	}
}
