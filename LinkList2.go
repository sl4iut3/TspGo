package main

//import "fmt"

type LinkList2 []Link

func (list *LinkList2) addLink(l Link) {
	*list = append(*list, l)
}

func (list LinkList2) getLink2(n1, n2 int) *Link {
	if n1 == n2 {
		return nil
	}
	if n1 > n2 {
		n1, n2 = n2, n1
	}
	for i, _ := range list {
		if list[i].n1 == n1 && list[i].n2 == n2 {
			return &list[i]
		}
	}
	return nil
}

func (list *LinkList2) getMinLink() Link {
	var minLink Link
	var minD int = 1e15
	var num int
	for i, l := range *list {
		if l.d < minD {
			minLink = l
			minD = l.d
			num = i
		}
	}
	//delete(list, minLink)
	//list=append(
	//	fmt.Println(num," ",len(*list))
	//if num==len(list)-1 {
	//	list=list[:num-1]
	//} else {
	*list = append((*list)[:num], (*list)[num+1:]...)
	//}
	//	fmt.Println(*list,len(*list))
	return minLink
}

func (list LinkList2) clone() LinkList2 {
	newList := LinkList2{}
	for _, l := range list {
		newList.addLink(l)
	}
	return newList
}

func (list LinkList2) toString() string {
	s := ""
	for i, _ := range list {
		s += "(" + list[i].toString() + ") "
	}
	return s
}

func createLinkList2(distances [][]int) LinkList2 {
	if len(distances) == 0 || len(distances) != len(distances[0]) {
		return nil
	}
	n := len(distances)
	var liste LinkList2
	//list := make(LinkList2,n*(n-1)/2)
	//list := make(LinkList2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			l := createLink(i, j, distances[i][j])
			liste.addLink(l)
			//liste= append(liste,l)
			//fmt.Println(liste)
			//list.addLink(createLink(i, j, distances[i][j]))
		}
	}

	return liste
}
