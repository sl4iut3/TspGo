package main

import ( 
	"fmt" 
)

type Link struct {
	n1, n2, d int
}

func (l Link) compareTo( l2 Link) (int) {
	return l.d-l2.d
}

func (l Link) String() (string) {
	return fmt.Sprintf("(%d, %d)[%d]", l.n1, l.n2, l.d)
}

func createLink(n1, n2, d int) (Link) {
	if n1 > n2 {
		return Link{n2, n1, d}
	} else {
		return Link{n1, n2, d}
	}
}
