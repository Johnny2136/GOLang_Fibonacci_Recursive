//Demo Fibonacci sequence generator
//Johnny2136
package main

import (
	"fmt"
	"time"
)

// fibonacci in GOlang gnited to 45 sequences
// fibo function returns recursive solution
func fiboRec(n int) int {
	if n < 2 {
		return n
	}
	return fiboRec(n-1) + fiboRec(n-2)
}

// main calls fiborec "i" number of times and prints to display
// the time.Since(start) gives elapsed time of main to execute
func main() {
	start := time.Now()

	for i := 0; i < 25; i++ {
		fmt.Println(fiboRec(i))
	}
	fmt.Println(time.Since(start))
}
