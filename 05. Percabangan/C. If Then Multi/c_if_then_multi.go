package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	if N > 0 && N % 2 == 0 {
		fmt.Println(N)
	}
}
