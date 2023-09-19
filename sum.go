package main

import (
	"fmt"
	"os"
	"strconv"
)

// func add(a int64, b int64) int64 {
// 	return (a + b)
	
// }

func main2(){
	args := os.Args

	a, err1 := strconv.ParseInt(args[1], 10, 64)
	b, err2 := strconv.ParseInt(args[2], 10, 64)

	if (err1 != nil || err2 != nil) {
		fmt.Println("Error", err1, err2)
	}

	fmt.Println("Sum:", a + b)

}
