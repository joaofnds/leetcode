package main

import "testing"

func TestLRUCache(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		actions  []string
		args     [][]int
		expected []any
	}{
		{
			name:     "Test1",
			capacity: 2,
			actions:  []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"},
			args:     [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}},
			expected: []any{nil, nil, nil, 1, nil, -1, nil, -1, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lruCache := Constructor(tt.capacity)
			for i, action := range tt.actions {
				switch action {
				case "LRUCache":
					continue
				case "put":
					lruCache.Put(tt.args[i][0], tt.args[i][1])
					if tt.expected[i] != nil {
						t.Errorf("expected %v but got nil", tt.expected[i])
					}
				case "get":
					result := lruCache.Get(tt.args[i][0])
					if result != tt.expected[i] {
						t.Errorf("expected %v but got %v", tt.expected[i], result)
					}
				default:
					t.Errorf("unknown action %s", action)
				}
			}
		})
	}
}
