package main

import (
	"fmt"
)

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

type taskList struct {
	tasks []*task
}

func (t *taskList) addToList(newTask *task) {
	t.tasks = append(t.tasks, newTask)
}

func (t *taskList) deleteOfList(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *taskList) printList() {
	for _, task := range t.tasks {
		fmt.Printf("Name: %s\n", task.name)
		fmt.Printf("Description: %s\n", task.description)
	}
}

func (t *taskList) printTaskCompleted() {
	for _, task := range t.tasks {
		if task.completed {
			fmt.Printf("Name: %s\n", task.name)
			fmt.Printf("Description: %s\n", task.description)
		}
	}
}

func main() {
	t1 := &task{
		name:        "Complete the go course",
		description: "Complete the Platzi go course of in this week",
	}

	t2 := &task{
		name:        "Complete the java course",
		description: "Complete the Platzi java course of in this week",
	}

	t3 := &task{
		name:        "Complete the python course",
		description: "Complete the Platzi python course of in this week",
	}

	t4 := &task{
		name:        "Complete the fury course",
		description: "Complete the Platzi fury course of in this week",
	}

	list := &taskList{
		tasks: []*task{
			t1, t2, t3,
		},
	}

	list.addToList(t4)

	fmt.Println("==========================")
	fmt.Println("========== for ===========")
	fmt.Println("==========================")

	for i := 0; i < len(list.tasks); i++ {
		fmt.Printf("Index: %v, Name: %v\n", i, list.tasks[i].name)
	}

	fmt.Println()
	fmt.Println("==========================")
	fmt.Println("======= for range ========")
	fmt.Println("==========================")

	for index, task := range list.tasks {
		fmt.Printf("Index: %v, Name: %v\n", index, task.name)
	}

	fmt.Println()
	fmt.Println("==========================")
	fmt.Println("======= use break ========")
	fmt.Println("==========================")
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println()
	fmt.Println("==========================")
	fmt.Println("===== use continue =======")
	fmt.Println("==========================")

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println()
	fmt.Println("==========================")
	fmt.Println("======= tasks list =======")
	fmt.Println("==========================")
	list.printList()

	fmt.Println()
	list.tasks[0].taskCompleted()
	list.tasks[1].taskCompleted()
	fmt.Println("==========================")
	fmt.Println("==== tasks completed =====")
	fmt.Println("==========================")
	list.printTaskCompleted()

	fmt.Println()
	fmt.Println("==========================")
	fmt.Println("========== maps ==========")
	fmt.Println("==========================")

	mapTasks := make(map[string]*taskList)
	mapTasks["Christian"] = list

	t5 := &task{
		name:        "Complete the Spring Boot course",
		description: "Complete the Platzi Spring Boot course of in this week",
	}

	t6 := &task{
		name:        "Complete the C# course",
		description: "Complete the Platzi C# course of in this week",
	}

	list2 := &taskList{
		tasks: []*task{
			t5, t6,
		},
	}

	mapTasks["Alberto"] = list2

	fmt.Println("Tareas de Christian")
	mapTasks["Christian"].printList()

	fmt.Println()
	fmt.Println("Tareas de Alberto")
	mapTasks["Alberto"].printList()
}
