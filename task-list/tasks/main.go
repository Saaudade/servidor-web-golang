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

type taskList struct {
	tasks []*task
}

func (t *taskList) addToList(newTask *task) {
	t.tasks = append(t.tasks, newTask)
}

func (t *taskList) deleteOfList(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
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

	for i := 0; i < len(list.tasks); i++ {
		fmt.Printf("Index: %v, Name: %v\n", i, list.tasks[i].name)
	}

	for index, task := range list.tasks {
		fmt.Printf("Index: %v, Name: %v\n", index, task.name)
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

}
