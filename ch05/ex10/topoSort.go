package main

import (
	"fmt"
)

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d: \t%s\n", i+1, course)
	}

	fmt.Printf("\ntopological Sort : %v\n", isTopoSorted(topoSort(prereqs)))
}

var prereqs = map[string]itemSet{
	"algorithms": newItemSet("data structures"),
	"calculus":   newItemSet("linear algebra"),
	"compilers": newItemSet(
		"data structures",
		"format languages",
		"computer organization",
	),
	"data structures":       newItemSet("discrate math"),
	"databases":             newItemSet("data structures"),
	"discrate math":         newItemSet("intro to programming"),
	"format languages":      newItemSet("discrate math"),
	"networks":              newItemSet("operating systems"),
	"operating systems":     newItemSet("data structures", "computer organization"),
	"programming languages": newItemSet("data structures", "computer organization"),
}

type itemSet map[string]bool

func (i itemSet) add(items ...string) {
	for _, item := range items {
		i[item] = true
	}
}

func newItemSet(items ...string) itemSet {
	is := make(itemSet)
	is.add(items...)

	return is
}

func topoSort(m map[string]itemSet) []string {
	var order []string
	seen := make(map[string]bool)

	var visitAll func(is itemSet)
	visitAll = func(is itemSet) {
		for key := range is {
			if !seen[key] {
				seen[key] = true
				visitAll(m[key])
				order = append(order, key)
			}
		}
	}

	/*
		var key []string
		for k := range m {
			key = append(key, k)
		}
		sort.Strings(key)
	*/
	is := newItemSet()
	for item := range m {
		is.add(item)
	}
	visitAll(is)

	return order
}

func isTopoSorted(order []string) bool {
	checkList := make(map[string]int)
	for i, item := range order {
		checkList[item] = i
	}

	for item, i := range checkList {
		for prereq := range prereqs[item] {
			if i < checkList[prereq] {
				return false
			}
		}
	}
	return true
}
