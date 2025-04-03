package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		inputs   []string
		expected [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		{[]string{"", ""}, [][]string{{"", ""}}},
		{[]string{"a"}, [][]string{{"a"}}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("groupAnagrams(%v)", test.inputs), func(t *testing.T) {
			actual := groupAnagrams(test.inputs)

			sortAnagramGroups(actual)
			sortAnagramGroups(test.expected)

			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("expected %q, got %q", test.expected, actual)
			}
		})
	}
}

func sortAnagramGroups(groups [][]string) {
	for _, group := range groups {
		sort.Strings(group)
	}

	sort.Slice(groups, func(i, j int) bool {
		if len(groups[i]) == 0 {
			return true
		}

		if len(groups[j]) == 0 {
			return false
		}

		return groups[i][0] < groups[j][0]
	})
}
