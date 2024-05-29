package main

import "fmt"

var a bool

func main() {
	fmt.Println(a)
	a = true
	fmt.Println(a)
	a = 10 == 100
	b := 10 == 100
	c := 10 > 100
	d := 10 < 100
	e := 10 <= 10
	f := 10 >= 10

	fmt.Println(a, b, c, d, e, f)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

}

/*
Operadores relacionais
== igual 
<= menor ou igual
>= maior ou igual
!= diferente
< menor que
> maior que
*/