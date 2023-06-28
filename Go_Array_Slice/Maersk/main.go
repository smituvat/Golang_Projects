/*Given a string, output occurrences of each character in order of most occurrences to fewest.
If two characters have the same amount of occurrences, output them in alphabetical order.
Example input: "ccaaaaabbddd"
Example output:
5 a
3 d
2 b
2 c
*/

package main

import (
	"fmt"
	"sort"
)

func main() {

	s := "ccaaaaabbddd"
	sortAlpha(s)

}

type Pair struct {
	Key   string
	Value int
}

func sortAlpha(s string) {

	m := make(map[string]int)

	for _, v := range s {
		s := string(v)
		m[s]++
	}

	// Convert the map to a slice of key-value pairs
	pairs := make([]Pair, 0, len(m))
	for k, v := range m {
		pairs = append(pairs, Pair{k, v})
	}

	// Sort the slice in decreasing order of values
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})

	// Reconstruct the sorted map
	sortedMap := make(map[string]int)
	for _, pair := range pairs {
		sortedMap[pair.Key] = pair.Value
	}

	// Print the sorted map
	for k, v := range sortedMap {
		fmt.Printf("%s: %d\n", k, v)
	}
}
