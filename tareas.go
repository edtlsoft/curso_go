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

type task struct {
	nombre string
	descripcion string
	completado bool
}

func (t *task) marcarCompleta() {
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

	//fmt.Println(lista.tasks[0])
	lista.addTaskInList(t3)
	//fmt.Println(lista.tasks[2])
	//fmt.Println(len(lista.tasks))

	//for i:=0; i < len(lista.tasks); i++ {
	//	fmt.Println("Indice: ", i, " nombre: ", lista.tasks[i].nombre)
	//}

	for indice, tarea := range lista.tasks {
		fmt.Println("Indice: ", indice, " nombre: ", tarea.nombre)
	}

	
}