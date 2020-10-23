package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	items, cs := topoSort(prereqs)
	for i, course := range items {
		fmt.Printf("%d: \t%s\n", i+1, course)
	}

	fmt.Printf("\n%s\n", strings.Repeat("-", 60))
	fmt.Printf("circulate coruse : %d\n", len(cs))
	for _, c := range cs {
		fmt.Printf("%s \n", c)
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
	"operating systems":     {"data structures", "computer organization", "networks"},
	"programming languages": {"data structures", "computer organization"},
	"linear algebra":        {"calculus"},
}

func topoSort(m map[string][]string) ([]string, []string) {
	var order []string
	var circulateCourse []string
	seen := make(map[string]bool)

	var visitAll func(items []string) []string
	visitAll = func(items []string) []string {
		var prereq []string

		for _, item := range items {
			prereq = append(prereq, item)

			if !seen[item] {
				seen[item] = true
				prereq = visitAll(m[item])
				order = append(order, item)

				for _, pre := range prereq {
					if item == pre {
						circulateCourse = append(circulateCourse, item)
					}
				}
			}
		}

		return prereq
	}

	var key []string
	for k := range m {
		key = append(key, k)
	}
	sort.Strings(key)
	visitAll(key)

	return order, circulateCourse
}
