package main

import (
	"fmt"
	"strings"
	"time"
)

func fr(leafs *[][]int, remaining []int, begin []int) {
	if(len(remaining) == 2) {
		tab := []int{remaining[0]}
		tab = append(tab, begin...)
		tab = append(tab, remaining[1])
		*leafs = append(*leafs, tab)
	} else {
		for i, v := range remaining {
			tab1 := make([]int, i)
			copy(tab1, remaining[:i])
			tab1 = append(tab1, remaining[i+1:]...)

			tab2 := make([]int, len(begin))
			copy(tab2, begin)
			tab2 = append(tab2, v)

			fr(leafs, tab1, tab2)
		}
	}
}

func frs(leafs *[]string, remaining string, begin string) {
	tab := strings.Split(remaining, " ")
	if len(tab) == 2 {
		*leafs = append(*leafs, tab[0] + " " + begin + " " + tab[1])
	} else {
		for i, v := range tab {
			frs(leafs, strings.Trim(strings.Join(tab[:i], " ") + " " + strings.Join(tab[i+1:], " "), " "), strings.Trim(begin + " " + v, " "))
		}
	}
}

func mainTree() {
	leafs := [][]int{}
	remaining := []int{1, 2, 3, 4}
	var tic, toc time.Time
	tic = time.Now()
	fr(&leafs, remaining, []int{})
	toc = time.Now()
	fmt.Println(toc.Sub(tic))
	fmt.Println(leafs)
}

func mainStringTree() {
	leafs := []string{}
	remaining := "1 2 3 4"
	var tic, toc time.Time
	tic = time.Now()
	frs(&leafs, remaining, "")
	toc = time.Now()
	fmt.Println(toc.Sub(tic))
	fmt.Println(strings.Join(leafs, " | "))
}

