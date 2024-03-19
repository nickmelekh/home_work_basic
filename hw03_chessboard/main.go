package main

import (
	"bytes"
	"fmt"
)

func main() {
	var size int
	var b bytes.Buffer

	fmt.Scanf("%d", &size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i%2 == 0 {
				if j%2 == 0 {
					b.WriteString("#")
				} else {
					b.WriteString(" ")
				}
			} else {
				if j%2 == 0 {
					b.WriteString(" ")
				} else {
					b.WriteString("#")
				}
			}

		}

		b.WriteString("\n")
	}

	fmt.Println(b.String())
}
