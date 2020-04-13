// Упражнение 4.5 - удалять смежные дубликаты не занимая доп. память
package main

import "fmt"

func main() {
	strings := []string{"one", "two", "two", "three", "four"}
	fmt.Println(dupl(strings))
}

func dupl(strings []string) []string {
	last := ""
	correction := 0

	for i, s := range strings {
		if s == last {
			correction++
		}
		strings[i-correction] = s
		last = s
		fmt.Println(i-correction, s)
	}
	return strings[:len(strings)-correction]
}
