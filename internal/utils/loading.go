package utils

import (
	"fmt"
	"time"
)

func Loading(done <-chan struct{}, promt string) {
	spinner := []rune{'|', '/', '-', '\\'}
	i := 0
	for {
		select {
		case <-done:
			return
		default:
			fmt.Printf("\r%s  %c", promt, spinner[i%len(spinner)])
			i++
			time.Sleep(100 * time.Millisecond)
		}

	}

}
