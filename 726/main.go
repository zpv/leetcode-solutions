package main

import (
	"fmt"
	"sort"
	"strconv"
	"unicode"
)

func countOfAtoms(formula string) string {
	m, _ := countAtomsHelper(formula, 0)
	atoms := make([]string, 0, len(m))
	for k := range m {
		atoms = append(atoms, k)
	}
	sort.Strings(atoms)

	output := ""

	for _, v := range atoms {
		if m[v] > 1 {
			output += fmt.Sprintf("%v%d", v, m[v])
		} else {
			output += v
		}
	}
	return output
}

func countAtomsHelper(formula string, cursor int) (map[string]int, int) {
	countMap := make(map[string]int)
	curElement := ""
	searching := false

	for cursor < len(formula) && rune(formula[cursor]) != rune(')') {
		if rune(formula[cursor]) == rune('(') {
			innerMap, newCursor := countAtomsHelper(formula, cursor+1)
			for k, v := range innerMap {
				countMap[k] += v
			}
			cursor = newCursor + 1
			continue
		}

		if unicode.IsNumber(rune(formula[cursor])) {
			if searching {
				curNumString := string(formula[cursor])
				for cursor+1 < len(formula) && unicode.IsNumber(rune(formula[cursor+1])) {
					cursor++
					curNumString += string(formula[cursor])
				}

				num, err := strconv.Atoi(curNumString)
				if err == nil {
					countMap[curElement] += num
					searching = false
					curElement = ""
				}
			}
		} else if unicode.IsUpper(rune(formula[cursor])) {
			if searching {
				countMap[curElement]++
				curElement = string(formula[cursor])
			} else {
				searching = true
				curElement += string(formula[cursor])
			}
		} else {
			curElement += string(formula[cursor])
		}
		cursor++
	}
	cursor++
	if curElement != "" {
		countMap[curElement]++
	}
	if cursor < len(formula) && unicode.IsNumber(rune(formula[cursor])) {
		curNumString := string(formula[cursor])
		for cursor+1 < len(formula) && unicode.IsNumber(rune(formula[cursor+1])) {
			cursor++
			curNumString += string(formula[cursor])
		}

		for k, v := range countMap {
			num, err := strconv.Atoi(curNumString)
			if err == nil {
				countMap[k] = v * num
			}

		}
	}

	return countMap, cursor
}

func main() {
	countOfAtoms("K69(ON(SO3)2)2")
}
