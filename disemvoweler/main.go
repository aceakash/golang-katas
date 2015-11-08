package main

import (
	"regexp"
	"strings"
)

func disemvowel(input string) []string {
	vowelsAndSpaces := regexp.MustCompile("[aeiou\\s]")
	noVowels := vowelsAndSpaces.ReplaceAllString(input, "")
	vowelsOnly := strings.Replace(strings.Join(vowelsAndSpaces.FindAllString(input, -1), ""), " ", "", -1)
	return []string{noVowels, vowelsOnly}
}

func main() {

}
