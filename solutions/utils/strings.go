package utils

import (
	"log"
	"strconv"
)

func Stoi(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return val
}
