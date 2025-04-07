package main

import "strings"

func simplifyPath(path string) string {
	splits := strings.Split(path, "/")
	stack := make([]string, 0, len(splits)/2)

	for _, part := range splits {
		switch part {
		case "", ".": // do nothing
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, part)
		}
	}

	return "/" + strings.Join(stack, "/")
}
