package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n:=0
	m:=1
	result:=0
	return func() int {
		result=n+m
		tmp:=n
		n=m
		m=m+tmp
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
