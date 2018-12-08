package main

import (
	"fmt"
	"strings"
)

func main() {
	r := strings.NewReader("Hello World!")
	b := make([]byte, 6)

	for {
		n, err := r.Read(b)
		if err != nil {
			fmt.Printf("got error: %v\n", err)
			break
		}
		fmt.Printf("read: %d\n", n)
		fmt.Printf("content:%s\n", string(b[:n]))
	}
}
