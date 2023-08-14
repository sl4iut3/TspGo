package main

import "fmt"

type Node struct {
	isLeafNode bool

	trip Trip

	link        Link
	left, right *Node
}

func printTabs(n int) (string) {
	s := ""
	for i := 0; i < n; i++ {
		s += " | "
	}
	return s
}

func (node Node) treeToString(depth int) (string) {
	if(node.isLeafNode) {
		return fmt.Sprintln(node.trip)
	} else {
		s := fmt.Sprintln(node.link)
		s += printTabs(depth)
		s += "<-- " + node.left.treeToString(depth + 1)
		s += printTabs(depth)
		return s + "--> " + node.right.treeToString(depth + 1)
	}
}

func (node Node) String() (string) {
	return node.treeToString(1)
}

func createTree(trips []Trip, links LinkList2) (*Node) {
//	fmt.Println("BEGIN")
//	fmt.Println(trips, len(trips))
//	fmt.Println(links, len(links))
	node := Node{}
	if len(trips) == 1 {
		node.isLeafNode = true
		node.trip = trips[0]
//		fmt.Println("END " + fmt.Sprint(node.trip))
		return &node
	} else {
		node.isLeafNode = false
	}
	for {
		node.link = links.getMinLink()
//		fmt.Println(node.link)
		leftTrips := []Trip{}
		rightTrips := []Trip{}
		for _, trip := range trips {
			if trip.hasLink(node.link) {
				leftTrips = append(leftTrips, trip)
			} else {
				rightTrips = append(rightTrips, trip)
			}
		}
//		fmt.Println("left ->", leftTrips, len(leftTrips))
//		fmt.Println("right ->", rightTrips, len(rightTrips))
		if len(leftTrips) > 0 && len(rightTrips) > 0 {
//			fmt.Println("LEFT")
			node.left = createTree(leftTrips, links.clone())
//			fmt.Println("RIGHT")
			node.right = createTree(rightTrips, links.clone())
			return &node
		}
	}
}

func testNode() {
	distances := [][]int{
		{ 0,  8, 39, 37, 50}, 
		{ 8,  0, 45, 47, 49}, 
		{39, 45,  0,  9, 21}, 
		{37, 47,  9,  0, 15}, 
		{50, 49, 21, 15,  0}}
	fmt.Println("Distances Matrix :")
	for i := 0; i < len(distances); i++ {
		fmt.Println(distances[i])
	}
	fmt.Println()

	list := createLinkList2(distances)
	fmt.Println("Links :")
	for i := 0; i < len(list); i++ {
		fmt.Println(" - " + fmt.Sprint(list[i]))
	}
	fmt.Println()
	
	trips := make([]Trip, halfFact(len(distances) - 1))
	for i := 0; i < halfFact(len(distances) - 1); i++ {
		trips[i] = createTrip(i, len(distances) - 1, list)
	}
	fmt.Println("Trips :")
	for i := 0; i < len(trips); i++ {
		fmt.Println(" - " + fmt.Sprint(trips[i]) + " = " + fmt.Sprint(trips[i].length()))
	}
	fmt.Println()

	tree := createTree(trips, list)
	fmt.Println("Tree :")
	fmt.Println(tree)
}
