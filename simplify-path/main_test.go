package main

import (
	"fmt"
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"/home/", "/home"},
		{"/home//foo/", "/home/foo"},
		{"/home/user/Documents/../Pictures", "/home/user/Pictures"},
		{"/../", "/"},
		{"/.../a/../b/c/../d/./", "/.../b/d"},
		{"/qSNIo///b/../BtIY/DrG/./MTWto/.///./h/../dz/KBIma/y/wgatd/fhrm///ePfTD/Hef/TVM/../mhZ/JR/.///vnZ/cRCQu", "/qSNIo/BtIY/DrG/MTWto/dz/KBIma/y/wgatd/fhrm/ePfTD/Hef/mhZ/JR/vnZ/cRCQu"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("simplifyPath(%q)", test.input), func(t *testing.T) {
			actual := simplifyPath(test.input)
			if actual != test.expected {
				t.Errorf("expected %q, got %q", test.expected, actual)
			}
		})
	}
}

func BenchmarkSimplifyPath(b *testing.B) {
	tests := []struct {
		input    string
		expected string
	}{
		{"/home/", "/home"},
		{"/home//foo/", "/home/foo"},
		{"/home/user/Documents/../Pictures", "/home/user/Pictures"},
		{"/../", "/"},
		{"/.../a/../b/c/../d/./", "/.../b/d"},
		{"/qSNIo///b/../BtIY/DrG/./MTWto/.///./h/../dz/KBIma/y/wgatd/fhrm///ePfTD/Hef/TVM/../mhZ/JR/.///vnZ/cRCQu", "/qSNIo/BtIY/DrG/MTWto/dz/KBIma/y/wgatd/fhrm/ePfTD/Hef/mhZ/JR/vnZ/cRCQu"},
	}

	for _, test := range tests {
		b.Run("simplifyPath", func(b *testing.B) {
			for b.Loop() {
				actual := simplifyPath(test.input)
				if actual != test.expected {
					b.Errorf("expected %q, got %q", test.expected, actual)
				}
			}
		})
	}
}
