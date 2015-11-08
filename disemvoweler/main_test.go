package main

import (
	"reflect"
	"testing"
)

var disemvowelTests = []struct {
	input      string
	consonants string
	vowels     string
}{
	{"hello", "hll", "eo"},
	{"world", "wrld", "o"},
	{"hello beautiful world", "hllbtflwrld", "eoeauiuo"},
	{"two drums and a cymbal fall off a cliff", "twdrmsndcymblfllffclff", "ouaaaaoai"},
	{"all those who believe in psychokinesis raise my hand", "llthswhblvnpsychknssrsmyhnd", "aoeoeieeioieiaiea"},
	{"did you hear about the excellent farmer who was outstanding in his field", "ddyhrbtthxcllntfrmrwhwststndngnhsfld", "ioueaaoueeeeaeoaouaiiiie"},
}

func TestDisemvowel(t *testing.T) {
	for _, tt := range disemvowelTests {
		actual, expected := disemvowel(tt.input), []string{tt.consonants, tt.vowels}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("disemvowel '%v' return %v, expected %v", tt.input, actual, expected)
		}
	}

}
