package utils

import (
	"strings"
	"log"
)

func Format(unformattedStr string) string {
	log.Print("string passed::", unformattedStr)
	t := strings.TrimSpace(unformattedStr)
	log.Print("string returned::", t)
	return t
}
