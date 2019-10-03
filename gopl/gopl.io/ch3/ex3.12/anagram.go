package anagram

import (
	"reflect"
)

func getCharCounts(s string, charmap map[rune]int) {
	for _, c := range s {
		charmap[c]++
	}
}

func IsAnagram(s1 string, s2 string) bool {
	map1 := make(map[rune]int)
	map2 := make(map[rune]int)
	getCharCounts(s1, map1)
	getCharCounts(s2, map2)

	return reflect.DeepEqual(map1, map2)

}
