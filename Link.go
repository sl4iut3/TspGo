package main

import (
	"fmt"
	"strconv"
)

type Link struct {
	n1, n2, d int
}

func (l Link) compareTo(l2 Link) int {
	return l.d - l2.d
}

func (l Link) toString() string {
	return "n1=" + strconv.Itoa(l.n1) +
		", n2=" + fmt.Sprint(l.n2) +
		", d=" + fmt.Sprint(l.d)
}

func createLink(n1, n2, d int) Link {
	if n1 > n2 {
		return Link{n2, n1, d}
	} else {
		return Link{n1, n2, d}
	}
}
