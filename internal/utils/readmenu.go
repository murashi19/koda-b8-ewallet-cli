package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// ReadMenuChoice membaca input user dengan aman.
// Mengembalikan -1 jika input tidak valid (bukan angka).
func ReadMenuChoice(prompt string) int {
	fmt.Print(prompt)
	line, err := reader.ReadString('\n')
	if err != nil {
		return -1
	}

	line = strings.TrimSpace(line)
	choice, err := strconv.Atoi(line)
	if err != nil {
		return -1
	}
	return choice
}
