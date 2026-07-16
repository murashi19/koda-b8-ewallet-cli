package utils

import (
	"bufio"
	"fmt"
	"os"
)

func EnterBack() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press Enter to back...")
	reader.ReadString('\n')
}
