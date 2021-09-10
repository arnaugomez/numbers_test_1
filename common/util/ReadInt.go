package util

import (
	"fmt"
	"strconv"
	"strings"
)

func ReadInt(question ...interface{}) int {
	fmt.Print(question...)
	fmt.Print(" ")
	for {
		readText := ReadSafe()
		text := strings.TrimSuffix(readText, "\n")
		i, err := strconv.Atoi(text)
		if err == nil {
			return i
		}
		fmt.Print("Please enter a number ")
	}
}
