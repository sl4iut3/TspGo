package main

import (
	"fmt"
	"time"
)

func rawHalfFact(n int) int {
	v := 1
	for i := 3; i <= n; i++ {
		v *= n
	}
	return v
}

func main() {
	var bufHalfFact = bufferize(rawHalfFact)
	var tic, toc time.Time
	tic = time.Now()
	for i := 0; i < 16; i++ {
		bufHalfFact(i)
	}
	toc = time.Now()
	fmt.Println(toc.Sub(tic))
	tic = time.Now()
	for i := 0; i < 16; i++ {
		bufHalfFact(i)
	}
	toc = time.Now()
	fmt.Println(toc.Sub(tic))
}
