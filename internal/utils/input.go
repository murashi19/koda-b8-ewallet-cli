package utils

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func Input(promt string) string {
	fmt.Print(promt)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		panic("Input dihentikan")
	}
	return scanner.Text()
}
