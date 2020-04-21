package main

import (
	"strings"
	"testing"
)

func TestHello(t *testing.T) {
	pattern := "/hello/hahaha/e"
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	t.Log(vs)
	for k, item := range vs {
		if item != "" {
			t.Log(k)
			t.Log(item)
			t.Log(item[0])
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	t.Log(parts)
}
