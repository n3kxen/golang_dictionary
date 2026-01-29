// Edgars Krauja 241RDB039
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func addSynonyms(synonyms map[string]map[string]bool, w1, w2 string) {
	if synonyms[w1] == nil {
		synonyms[w1] = make(map[string]bool)
	}
	if synonyms[w2] == nil {
		synonyms[w2] = make(map[string]bool)
	}

	synonyms[w1][w2] = true
	synonyms[w2][w1] = true
}

func countSynonyms(synonyms map[string]map[string]bool, word string) int {
	if synonyms[word] != nil {
		return len(synonyms[word])
	}
	return 0
}

func checkSynonyms(synonyms map[string]map[string]bool, w1, w2 string) bool {
	if synonyms[w1] != nil && synonyms[w1][w2] {
		return true
	}
	return false
}

func main() {
	synonyms := make(map[string]map[string]bool)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		parts := strings.Fields(input)

		if len(parts) == 0 {
			continue
		}

		cmd := parts[0]

		switch cmd {
		case "add":
			if len(parts) < 3 {
				continue
			}
			w1 := parts[1]
			w2 := parts[2]
			addSynonyms(synonyms, w1, w2)

		case "count":
			if len(parts) < 2 {
				continue
			}
			word := parts[1]
			count := countSynonyms(synonyms, word)
			fmt.Println(count)

		case "check":
			if len(parts) < 3 {
				continue
			}
			w1 := parts[1]
			w2 := parts[2]
			if checkSynonyms(synonyms, w1, w2) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}

		case "done":
			return
		}
	}
}
