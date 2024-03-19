package main

import (
	"bytes"
	"fmt"
)

var (
	size int
	b    bytes.Buffer
)

func main() {
	fmt.Scanf("%d", &size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (i+j)%2 == 0 {
				b.WriteString("#")
			} else {
				b.WriteString(" ")
			}
		}

		b.WriteString("\n")
	}

	fmt.Println(b.String())
}
