/*
Exercise 4.9: Write a program wordfreq to report the frequency of each word in an input text ﬁle.
Call input.Split(bufio.ScanWords) before the ﬁrst call to Scan to break the input into words instead of lines
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s text_file_name\n", os.Args[0])
		return
	}
	fmt.Printf("%s\n%s\n", os.Args[0], os.Args[1])
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error open file %s, %v\n", os.Args[1], err)
		return
	}

	input := bufio.NewScanner(file)
	freqMap := make(map[string]int)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		wordFreq(input.Text(), freqMap)
	}

	for k, v := range freqMap {
		fmt.Printf("%s,%d\n", k, v)
	}
}

func wordFreq(word string, freqMap map[string]int) {
	freqMap[word]++
}
