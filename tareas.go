package main

import (
	"fmt"
)

type taskList struct {
	tasks []*task
}

func (tl *taskList) addTaskInList(t *task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *taskList) removeTaskOfList(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index + 1:]...)
}

func (tl *taskList) printList() {
	for _, task := range tl.tasks {
		fmt.Println("Nombre:", task.nombre)
		fmt.Println("Descripcion:", task.descripcion)
	}
}

func (tl *taskList) printListCompleted() {
	for _, task := range tl.tasks {
		if task.completado {
			fmt.Println("Nombre:", task.nombre)
			fmt.Println("Descripcion:", task.descripcion)
		}
	}
}


type task struct {
	nombre string
	descripcion string
	completado bool
}

func (t *task) marcarCompletada() {
	t.completado = true
}

func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	t1 := &task{
		nombre: "Aprender Go",
		descripcion: "Completar el curos de Go esta semana.",
	}

	t2 := &task{
		nombre: "Aprender Python",
		descripcion: "Completar el curos de Python esta semana.",
	}

	t3 := &task{
		nombre: "Aprender NodeJS",
		descripcion: "Completar el curos de NodeJS esta semana.",
	}

	fmt.Println(t1) // &{Aprender Go Completar el curos de Go esta semana. false}

	lista := &taskList {
		tasks: [] *task {
			t1, t2,
		},
	}

	lista.addTaskInList(t3)
	lista.printList();

	lista.tasks[0].marcarCompletada()
	fmt.Println("I M P R I M I R  L I S T A  D E  T A R E A S  C O M P L E T A D A S")
	lista.printListCompleted();
	
}