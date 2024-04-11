package main

import (
	"strings"
)

func ParseStep1Result(result string) (string, string) {
	splited1 := strings.Split(result, `input[name=checksum]').value`)
	checksum := strings.Split(splited1[1], `'`)[1]

	splited := strings.Split(result, `<span  format="24h" class="timer" counting="up" value="`)
	timestamp := strings.Split(splited[1], `"`)[0]

	return checksum, timestamp
}
