package main

import (
	"fmt"
	"time"
	"net"
)

func rawHalfFact(n int) int {
	v := 1
	for i := 3; i <= n; i++ {
		v *= n
	}
	return v
}

func mult2 (n int) int { return  2*n }

func codeDemain() {
	var bufHalfFact = bufferize(rawHalfFact)
	var m2 = bufferize(mult2)
	var tic, toc time.Time
	tic = time.Now()
	p := func() { for i := 0; i < 16; i++ {
		bufHalfFact(i) 
		m2(i)
	}}
	p()
	toc = time.Now()
	fmt.Println(toc.Sub(tic))
	tic = time.Now()
	for i := 0; i < 16; i++ {
		bufHalfFact(i); m2(i)
	}
	toc = time.Now()
	fmt.Println(bufHalfFact(5),"  ",m2(5))
	fmt.Println(toc.Sub(tic))

	fmt.Println(net.Interfaces())
	fmt.Println(net.LookupAddr("8.8.8.8"))
}

func main3() {
	//p := codeDemain
	//p()
	t := []int{1,2,3,4,5,6}
	var i int
	for i=1; i<len(t)-1; i++ {
		var subt []int
		subt = append(subt,t[:i]...)
		subt = append(subt,t[i+1:]...)
		fmt.Println(i," ",subt," ",t[:i],t[i+1:])
	}
}

