// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	//	for i := 1; i < len(os.Args); i++ {
	//		fmt.Printf("  %s\n", comma(os.Args[i]))
	//	}
	for input.Scan() {
		fmt.Printf("  %s\n", comma(input.Text()))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	for i := 0; i < len(s); i++ {
		buf.WriteByte(s[i])
		if (i+1)%3 == 0 {
			buf.WriteString(",")
		}
	}

	return buf.String()
}

//!-
