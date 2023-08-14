package main

import "fmt"

type Node struct {
	isLeafNode bool

	trip Trip

	link        Link
	left, right *Node
}

func (node Node) treeToString(depth int) string {
	//if(node.isLeafNode) {
	s := ""
	for i := 0; i < depth; i++ {
		s += "\t"
	}
	if node.isLeafNode {
		return s + node.trip.toString() + "\n"
	} else {
		//	s := ""
		//	for i := 0; i < depth; i++ {
		//		s += "\t"
		//	}
		s += node.link.toString() + "\n"
		s += "<-" + node.left.treeToString(depth+1)
		return s + "->" + node.right.treeToString(depth+1)
	}
}

func (node Node) toString() string {
	return node.treeToString(0)
}

func createTree(trips []Trip, links LinkList2) *Node {
	node := Node{}
	if len(trips) == 1 {
		node.isLeafNode = true
		node.trip = trips[0]
		fmt.Println(node.trip)
		return &node
	} else {
		node.isLeafNode = false
	}
	for {
		node.link = links.getMinLink()
		fmt.Println(node.link.toString())
		leftTrips := []Trip{}
		rightTrips := []Trip{}
		for _, trip := range trips {
			if trip.hasLink(node.link) {
				leftTrips = append(leftTrips, trip)
			} else {
				rightTrips = append(rightTrips, trip)
			}
		}
		fmt.Println("left ->", leftTrips, len(leftTrips))
		fmt.Println("right ->", rightTrips, len(rightTrips))
		if len(leftTrips) > 0 && len(rightTrips) > 0 {
			node.left = createTree(leftTrips, links.clone())
			node.right = createTree(rightTrips, links.clone())
			return &node
		}
	}
}
