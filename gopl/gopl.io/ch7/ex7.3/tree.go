package main

import (
	"bytes"
	"fmt"
	"strconv"
	//"strconv"
)

//!+
type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func (tr *tree) Sort(values []int) *tree {
	for _, v := range values {
		//root = tr.add(root, v)
		tr = tr.add(v)
	}
	tr.appendValues(values[:0], tr)
	return tr
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
// sort in place (in the orginal array)
func (tr *tree) appendValues(values []int, t *tree) []int {
	if t != nil {
		values = tr.appendValues(values, t.left)
		values = append(values, t.value)
		values = tr.appendValues(values, t.right)
	}
	return values
}

func (tr *tree) add(value int) *tree {
	if tr == nil {
		// Equivalent to return &tree{value: value}.
		tr = new(tree)
		tr.value = value
		return tr
	}
	if value < tr.value {
		tr.left = tr.left.add(value)
	} else {
		tr.right = tr.right.add(value)
	}
	return tr
}

func (tr *tree) String() string {
	buffer := new(bytes.Buffer)
	tr.traverse(buffer)
	return buffer.String()
}

func (tr *tree) traverse(buffer *bytes.Buffer) *bytes.Buffer {
	if tr != nil {
		buffer = tr.left.traverse(buffer)
		buffer.WriteString(strconv.Itoa(tr.value))
		buffer.WriteString(",")
		buffer = tr.right.traverse(buffer)
	}
	return buffer
}

func (tr *tree) Print() {
	//fmt.Println("Start to print........\n")
	if tr != nil {
		tr.left.Print()
		fmt.Print(tr.value, " ")
		tr.right.Print()
	}
}

//!-

func main() {
	var values = []int{3, 4, 1, 9, 0, 2, 11}
	var tr *tree
	tr = tr.Sort(values)
	//tr.Print()
	//str := tr.String()
	fmt.Println(tr)
	//fmt.Println(values)

}
