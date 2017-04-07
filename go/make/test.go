package main

import (
	"fmt"
)

func main() {

	m := make([]byte, 100)

	for i, j := range m {
		fmt.Printf("%v %v\n", i, j)
	}

}
