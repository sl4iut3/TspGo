package main

import "fmt"

type NodeList []bool

func (l NodeList) get(v int) int {
//	fmt.Print("v=", v)
//	if l[v] == false {
//		l[v] = true
//		return v + 1
//	} else {
		for i := 0; i < len(l); i++ {
			if l[i] {
				continue
			} else if v > 0 {
				v--
			} else {
				l[i] = true
//				fmt.Println(" ", i+1, " l=", l)
				return i + 1
			}
		}
	//}
	return -1
}

func createNodeList(n int) NodeList {
	l := make([]bool, n)
	/*fmt.Println("**** l=",l)
	for i := 0; i < n; i++ {
		l[i] = false
	}*/
	return l
}

type Trip []Link

func (trip *Trip) addLink(link Link) {
	*trip = append(*trip, link)
}

func (trip Trip) hasLink(link Link) bool {
	for _, l := range trip {
		if l.n1 == link.n1 && l.n2 == link.n2 {
			return true
		}
	}
	return false
}

func (trip Trip) toString() string {
	s := "[ "
	v := trip[0].n2
	for i, l := range trip {
		if i < len(trip) {
			if l.n1 == v {
				s += fmt.Sprint(l.n2) + " "
				v = l.n2
			} else {
				s += fmt.Sprint(l.n1) + " "
				v = l.n1
			}
		}
	}
	return s + "]"
}

func bufferize(f func(int) int) func(int) int {
	var buffer [20]int
	return func(n int) int {
		if buffer[n] == 0 {
			buffer[n] = f(n)
		}
		return buffer[n]
	}
}

var halfFact = buildHalfFact()

// func halfFact(n int) int {
func buildHalfFact() func(int) int {
	var hf [20]int
	return func(n int) int {
		if hf[n] != 0 {
			return hf[n]
		} else {
//			fmt.Println("calcul ", n)
			v := 1
			for i := 3; i <= n; i++ {
				v *= i
			}
			hf[n] = v
			return v
		}
	}
}

func createTrip(v int, n int, list LinkList2) Trip {
	visitedNodes := createNodeList(n)
	trip := make([]int, n)

	for i := 1; i < n-1; i++ {
		trip[i] = visitedNodes.get(v / halfFact(n-i))
		v %= halfFact(n - i)
	}

	trip[0] = visitedNodes.get(0)
	trip[n-1] = visitedNodes.get(0)

//	return createTripFromArray(trip, list)
	return Trip{}
}

func createTripFromArray(t []int, list LinkList2) Trip {
	fmt.Println(t)
	trip := Trip{}
	trip.addLink(*(list.getLink2(0, t[0])))
	for i := 1; i < len(t); i++ {
		trip.addLink(*(list.getLink2(t[i-1], t[i])))
	}
	trip.addLink(*(list.getLink2(t[len(t)-1], 0)))
	return trip
}

func testHalfFact() {
	for i := 0; i < 16; i++ {
		fmt.Println(i, "\t", halfFact(i))
	}
}
