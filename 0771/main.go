package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("% x", numJewelsInStones("aA", "aAAbbbb"))
}

func numJewelsInStones(J string, S string) (count int) {
	for _, runeValue := range S {
		if strings.ContainsRune(J, runeValue) {
			count++
		}
	}

	return
}
