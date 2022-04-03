package main

import "fmt"

type task struct {
	name        string
	description string
	completed   bool
}

func (t *task) taskCompleted() {
	t.completed = true
}

func (t *task) setDescription(description string) {
	t.description = description
}

func (t *task) setName(name string) {
	t.name = name
}

func main() {
	t := &task{
		name:        "Complete the go course",
		description: "Complete the Platzi go course of in this week",
	}
	t.taskCompleted()
	t.setName("Completed the go course x2")
	t.setDescription("Completed the course very quickly")
	fmt.Printf("%+v\n", t)
}
