package util

import (
	"bufio"
	"fmt"
	"os"
)

func ReadSafe() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		text, err := reader.ReadString('\n')

		if err == nil || len(text) == 0 {
			return text
		}
	}
}
