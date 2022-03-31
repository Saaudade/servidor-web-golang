package main

import "fmt"

func main() {
	x := 25
	fmt.Println(x)
	setValue(x)
	fmt.Println(&x)
	y := &x
	fmt.Println(y)
	fmt.Println(*y)
}

func setValue(value int) {
	value = 36
}
