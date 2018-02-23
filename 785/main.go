package main

import (
	"fmt"
)

func main() {
	graph := [][]int{{}, {2}, {1}, {}, {}, {7, 8}, {7, 8, 9}, {5, 6}, {5, 6}, {6}, {12, 13, 14}, {12}, {10, 11}, {10}, {10}, {18}, {17, 18}, {16}, {15, 16}, {}, {22, 23, 24}, {22, 23, 24}, {20, 21}, {20, 21}, {20, 21}, {27, 28, 29}, {27, 28, 29}, {25, 26}, {25, 26}, {25, 26}, {32, 33, 34}, {33}, {30}, {30, 31}, {30}, {37, 39}, {38}, {35}, {36}, {35}, {44}, {43, 44}, {}, {41}, {40, 41}, {47, 48, 49}, {47, 48, 49}, {45, 46}, {45, 46}, {45, 46}}
	fmt.Println(isBipartite(graph))
}

func isBipartite(graph [][]int) bool {
	set := make(map[int]int)

	for i := 0; i < len(graph); i++ {
		if set[i] == 0 {
			top := 0
			stack := make([]int, 100000)
			stack[top] = i
			set[i] = 1

			for top >= 0 {
				node := stack[top]
				top--

				for _, nei := range graph[node] {
					if set[nei] == 0 {
						top++
						stack[top] = nei

						if set[node] == 1 {
							set[nei] = 2
						} else {
							set[nei] = 1
						}
					} else if set[nei] == set[node] {
						return false
					}
				}
			}
		}
	}
	return true
}
