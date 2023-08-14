/*
tsp : a command for simulate a new approach to TSP

Usage :

	tsp
*/
package main

import (
	"fmt"
	"time"
)

// println
// ... to avoid fmt.Println
// Rque: non visible ds go doc --all car non exportée (pas de majuscule)
func println[T string | int](s T) {
	fmt.Println(s)
}

//func main() { mainStringTree() }
func main() { mainTSP() }

func mainTSP() {
	fmt.Println("TSP!")

	distances := [][]int{
		{0, 1, 2, 3, 4},
		{10, 0, 5, 6, 7},
		{9, 8, 0, 8, 9},
		{7, 6, 5, 0, 10},
		{4, 3, 2, 1, 0}}

	for i := 0; i < len(distances); i++ {
		fmt.Println(distances[i])
	}

	list := createLinkList2(distances)
	fmt.Println(list)

	var tic, toc time.Time
	tic = time.Now()
	trips := make([]Trip, halfFact(len(distances)-1))
	for i := 0; i < halfFact(len(distances)-1); i++ {
		trips[i] = createTrip(i, len(distances)-1, list)
	}
	toc = time.Now()
	fmt.Println(toc.Sub(tic))

	tree := createTree(trips, list)
	fmt.Println(tree.toString())

}
