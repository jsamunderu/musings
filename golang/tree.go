package main

import "fmt"
import "strings"
import "strconv"

func TreeConstructor(strArr []string) string {
	graph := make([][]int, 0)
	for _, v := range strArr {
		tmp := strings.ReplaceAll(v, "(", "")
		tmp = strings.ReplaceAll(tmp, ")", "")
		tuple := strings.Split(tmp, ",")
		first, _ := strconv.Atoi(tuple[0])
		second, _ := strconv.Atoi(tuple[1])
		fmt.Println("first: ", first, " second ", second)
		if i := find(graph[:], second); i == -1 {
			entry := make([]int, 0)
			entry = append(entry, second)
			entry = append(entry, first)
			graph = append(graph, entry)
		} else {
			graph[i] = append(graph[i], first)
		}
	}
	m := make(map[int]bool)
	for _, v := range graph {
		fmt.Println(v[0])
		m[v[0]] = true
	}
	for _, v := range graph {
		disjoint := true
		for i := 1; i < len(v); i++ {
			if _, ok := m[v[i]]; ok {
				disjoint = false
				break
			}
		}
		if disjoint {
			return "false"
		}
	}
	return "true"

}

func find(a [][]int, k int) int {
	for i, v := range a {
		if v[0] == k {
			return i
		}
	}
	return -1
}
func main() {
	input := [...]string{"(1,2)", "(2,4)", "(7,2)"}
	//input := [...]string{"(1,2)", "(3,2)", "(2,12)", "(5,2)"}
	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(TreeConstructor(input[:]))

}
