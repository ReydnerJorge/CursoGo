package main

import "fmt"

func main() {
	x := 0
	y := 0.0
	z := false
	w := ""
	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y)
	fmt.Printf("%v, %T\n", z, z)
	fmt.Printf("%v, %T", w, w)
}

/*
Os zeros para cada tipo:
ints: 0
floats: 0.0
booleans: false
string: ""
pointers, functions, interfaces, slices, channels, maps: nil
*/