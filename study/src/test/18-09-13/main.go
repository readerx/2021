package main

import "fmt"

func main() {
	bs := []byte(nil)

	fmt.Printf("len: %d\n", len(bs))

	for _, x := range bs {
		fmt.Printf("len: %v\n", x)
	}
}
