package main

import "fmt"

func fibo_rec(n, u, u_1 int) (int) {
	if n == 0 {
		return 0
	} else if n > 2 {
		return fibo_rec(n - 1, u + u_1, u)
	} else {
		return u;
	}
}

// Fonction de calcul de la suite de Fibonacci
func Fibo(n int) (int) {
	return fibo_rec(n, 1, 1)
}

func testFibo() {
	for i := 0; i < 32; i++ {
		fmt.Println(fmt.Sprint(i) + "\t" + fmt.Sprint(Fibo(i)))
	}
}
