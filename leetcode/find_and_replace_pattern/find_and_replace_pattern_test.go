package findandreplacepattern

import (
	"reflect"
	"testing"
)

func TestFindAndReplacePattern(t *testing.T) {
	words := []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}
	pattern := "abb"
	expected := []string{"mee", "aqq"}
	result := findAndReplacePattern(words, pattern)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
