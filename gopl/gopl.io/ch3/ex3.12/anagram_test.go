package anagram

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	if IsAnagram("listen", "silent") != true {
		t.Error(`"listen", "silent"`)

	}
	if IsAnagram("test", "ttew") != false {
		t.Error(`"test", "ttew"`)

	}
	if IsAnagram("geeksforgeeks", "forgeeksgeeks") != true {
		t.Error(`"geeksforgeeks", "forgeeksgeeks"`)

	}
	if IsAnagram("triangle", "integral") != true {
		t.Error(`"triangle", "integral"`)

	}
	if IsAnagram("abd", "acb") != false {
		t.Error(`"abd", "acb"`)

	}

	if IsAnagram("中国", "国中") != true {
		t.Error(`"中国", "国中"`)

	}
}
