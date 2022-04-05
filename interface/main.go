package main

import "fmt"

type dog struct{}

type fish struct{}

type bird struct{}

func (dog) walk() string {
	return "I'm a dog and I walk"
}

func (fish) swim() string {
	return "I'm a fish and I swimming"
}

func (bird) fly() string {
	return "I'm a bird and I flying"
}

func dogMove(d dog) {
	fmt.Println(d.walk())
}

func fishMove(f fish) {
	fmt.Println(f.swim())
}

func birdMove(b bird) {
	fmt.Println(b.fly())
}

func main() {
	d := dog{}
	dogMove(d)
	f := fish{}
	fishMove(f)
	b := bird{}
	birdMove(b)
}
