package main

type LinkList map[Link]int

func (list LinkList) addLink(l Link) {
	list[l] = l.d
}

func (list LinkList) getLink(n1, n2 int) *Link {
	if n1 == n2 {
		return nil
	}
	if n1 > n2 {
		n1, n2 = n2, n1
	}
	for l, _ := range list {
		if l.n1 == n1 && l.n2 == n2 {
			return &l
		}
	}
	return nil
}

func (list LinkList) getMinLink() Link {
	var minLink Link
	var minD int = 1e15
	for link, d := range list {
		if d < minD {
			minLink = link
			minD = d
		}
	}
	delete(list, minLink)
	return minLink
}

func (list LinkList) toString() string {
	s := ""
	for l, _ := range list {
		s += "(" + l.toString() + ") "
	}
	return s
}

func createLinkList(distances [][]int) LinkList {
	if len(distances) == 0 || len(distances) != len(distances[0]) {
		return nil
	}

	list := LinkList{}
	for i := 0; i < len(distances); i++ {
		for j := i + 1; j < len(distances); j++ {
			list.addLink(createLink(i, j, distances[i][j]))
		}
	}

	return list
}
