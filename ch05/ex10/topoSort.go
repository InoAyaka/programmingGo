package main

import (
	"fmt"
	"sort"
)

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d: \t%s\n", i+1, course)
	}
}

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"format languages",
		"computer organization",
	},
	"data structures":       {"discrate math"},
	"databases":             {"data structures"},
	"discrate math":         {"intro to programming"},
	"format languages":      {"discrate math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)

	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var key []string
	for k := range m {
		key = append(key, k)
	}
	sort.Strings(key)
	visitAll(key)

	return order
}
