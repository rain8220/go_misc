//Exercise 4.5: Write anin-place function to eliminate adjacent duplicates in a []string slice
package main

import "fmt"

func main() {
	strs := []string{"aa", "aa", "cd", "de", "de", "de"}
	strs1 := eliminateDups(strs)
	fmt.Printf("%s \n", strs1)
}

func eliminateDups(s []string) []string {
	if len(s) < 2 {
		return s
	}
	c := s[0]
	k := 1
	for i := 1; i < len(s); i++ {
		if c != s[i] {
			s[k] = s[i]
			k++
			c = s[i]
		}
	}
	s = s[:k]
	return s

}
