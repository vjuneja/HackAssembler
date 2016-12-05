package main

import (
	"strings"
	"unicode"
)

func scrub(instruction string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
    		return -1
  		}
  		return r
	}, strings.Split(instruction, "//")[0]);
}
