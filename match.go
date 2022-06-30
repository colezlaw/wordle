package main

import (
	"errors"
	"strings"
)

var ErrUnequalLength = errors.New("lengths of word and state are unequal")
var ErrInvalidChar = errors.New("invalid state character")

// IsMatch returns whether word matches the
// state, where each character of state is one of:
//
// [A-Z] - a green letter in the Wordle state
// [a-z] - a yellow letter in the Wordle state
// _ - a blank in the Wordle state
// IsMatch will return an error if the lengths
// of state and word are not equal
func IsMatch(state, word, deny string) (bool, error) {
	for _, r := range []rune(strings.ToUpper(deny)) {
		if strings.ContainsRune(strings.ToUpper(word), r) {
			return false, nil
		}
	}
	if len(state) != len(word) {
		return false, ErrUnequalLength
	}

	w := []rune(strings.ToUpper(word))
	s := []rune(state)

	for i, l := range s {
		if l >= 'A' && l <= 'Z' {
			if w[i] != l {
				return false, nil
			}
			continue
		}
		if l >= 'a' && l <= 'z' {
			if w[i] == l-' ' {
				return false, nil
			}
			if !strings.ContainsRune(string(w), l-' ') {
				return false, nil
			}
			continue
		}
		if l != '_' {
			return false, ErrInvalidChar
		}
	}

	return true, nil
}
