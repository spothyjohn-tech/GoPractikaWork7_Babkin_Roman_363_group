package main

import (
	"fmt"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Status      string
	IsComplete  bool
}

type TaskManager struct {
	ID    int
	tasks []Task
}

func (taskManager *TaskManager) AddTask(title, description, status string) {
	newTask := Task{
		ID:          len(taskManager.tasks) + 1,
		Title:       title,
		Description: description,
		Status:      status,
		IsComplete:  false,
	}
	taskManager.tasks = append(taskManager.tasks, newTask)
	fmt.Println("Добавлена задача:", title)
}

func (taskManager *TaskManager) PrintAllTasks() {
	for _, task := range taskManager.tasks {
		fmt.Println("--------------------------------------------------")
		fmt.Println("Задача с ID:", task.ID, "имеет название:", task.Title)
		fmt.Println("Описание", task.Description, "Статус:", task.Status)
		fmt.Println("Выполнена:", task.IsComplete)
		fmt.Println("--------------------------------------------------")
	}
}

func (taskManager *TaskManager) DeleteTask() {
	for _, task := range taskManager.tasks {
		fmt.Println("ID:", task.ID, "Название задачи:", task.Title)
	}
	fmt.Println("Введите ID задачи для удаления:")
	var choice int
	fmt.Scanln(&choice)
	for i, task := range taskManager.tasks {
		if choice == task.ID {
			fmt.Println("Задача:", task.Title, "была удалена")
			taskManager.tasks = append(taskManager.tasks[:i], taskManager.tasks[i+1:]...)
			return
		}
	}
	fmt.Println("Задача с таким ID не найдена.")
}
func (taskManager *TaskManager) CompleteTask() {
	for _, task := range taskManager.tasks {
		fmt.Println("ID:", task.ID, "Название задачи:", task.Title)
	}
	fmt.Println("Введите ID задачи которую вы выполнили:")
	var choice int
	fmt.Scanln(&choice)
	for i, task := range taskManager.tasks {
		if choice == task.ID {
			fmt.Println("Задача с ID:", task.ID, "и названием:", task.Title, "отмечена как выполненная")
			taskManager.tasks[i].IsComplete = true
			return
		}
	}
	fmt.Println("Задача с таким ID не найдена.")
}
func (taskManager *TaskManager) PrintTasks() {
	fmt.Println("--------------------------------------------------")
	fmt.Println("По какому признаку вывести задачи:")
	fmt.Println("1. По статусу")
	fmt.Println("2. Выполненные")
	fmt.Println("3. Не выполненные")
	fmt.Println("4. По названию")
	fmt.Println("5. По ID")
	fmt.Println("--------------------------------------------------")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Println("Введите статус:")
		var status string
		fmt.Scanln(&status)
		for _, task := range taskManager.tasks {
			if status == task.Status {
				fmt.Println("--------------------------------------------------")
				fmt.Println("Задача с ID:", task.ID, "имеет название:", task.Title)
				fmt.Println("Описание", task.Description, "Статус:", task.Status)
				fmt.Println("Выполнена:", task.IsComplete)
				fmt.Println("--------------------------------------------------")
			}
		}
		break
	case 2:
		fmt.Println("Выполненные задачи:")
		for _, task := range taskManager.tasks {
			if task.IsComplete {
				fmt.Println("--------------------------------------------------")
				fmt.Println("Задача с ID:", task.ID, "имеет название:", task.Title)
				fmt.Println("Описание", task.Description, "Статус:", task.Status)
				fmt.Println("Выполнена:", task.IsComplete)
				fmt.Println("--------------------------------------------------")
			}
		}
		break
	case 3:
		fmt.Println("Не выполненные задачи:")
		for _, task := range taskManager.tasks {
			if !task.IsComplete {
				fmt.Println("--------------------------------------------------")
				fmt.Println("Задача с ID:", task.ID, "имеет название:", task.Title)
				fmt.Println("Описание", task.Description, "Статус:", task.Status)
				fmt.Println("Выполнена:", task.IsComplete)
				fmt.Println("--------------------------------------------------")
			}
		}
		break
	case 4:
		fmt.Println("Введите название:")
		var Title string
		fmt.Scanln(&Title)
		for _, task := range taskManager.tasks {
			if Title == task.Title {
				fmt.Println("--------------------------------------------------")
				fmt.Println("Задача с ID:", task.ID, "имеет название:", task.Title)
				fmt.Println("Описание", task.Description, "Статус:", task.Status)
				fmt.Println("Выполнена:", task.IsComplete)
				fmt.Println("--------------------------------------------------")
			}
		}
	case 5:
		var ID int
		fmt.Scanln(&ID)
		for _, task := range taskManager.tasks {
			if ID == task.ID {
				fmt.Println("--------------------------------------------------")
				fmt.Println("Задача с ID:", task.ID, "имеет название:", task.Title)
				fmt.Println("Описание", task.Description, "Статус:", task.Status)
				fmt.Println("Выполнена:", task.IsComplete)
				fmt.Println("--------------------------------------------------")
			}
		}
		break
	default:
		fmt.Println("Такого признака нет")
	}
}

func main() {
	var taskManager TaskManager
	taskManager.AddTask("Купить хлеб", "Зайти в магазин", "важно")
	taskManager.AddTask("Позвонить другу", "Узнать, как дела", "личное")
	taskManager.AddTask("Доделать домашку", "Это только седьмая задача по go...", "учёба")
	taskManager.PrintTasks()
	taskManager.CompleteTask()
	taskManager.DeleteTask()
	taskManager.PrintTasks()
}
