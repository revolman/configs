// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
	"log"
	"os"
	"sort"
)

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

//!-table

//!+main
func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string            // результирующий срез
	var parent string             // родительский ключ для проверки на циклы
	seen := make(map[string]bool) // проверка повторяемости

	var visitAll func(items []string) // функция сортировки
	visitAll = func(items []string) {
		for _, item := range items { // получаю ключи и для каждого ключа делаю действие
			if !seen[item] {
				seen[item] = true
				for _, dep := range m[item] { // проверяю все зависимости данного ключа на совпадение с родительским ключом
					if dep == parent { // если совпало, тогда это зацикливание и результат будет не очень корректным
						log.Fatalf("Обнаружен цикл мужду %s и %s!\n", dep, item)
						os.Exit(1)
					}
				}
				visitAll(m[item])
				order = append(order, item)
			}
			parent = item
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order
}

//!-main
