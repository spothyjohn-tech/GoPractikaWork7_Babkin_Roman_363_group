package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)


type Book struct {
	ID        int
	Title     string
	Author    string
	IsIssued  bool 
}

type Library struct {
	ID    int
	books []Book
}

func (lib *Library) AddBook(title string, author string) {
	newBook := Book{
		ID:       len(lib.books) + 1,
		Title:    title,
		Author:   author,
		IsIssued: false,
	}
	lib.books = append(lib.books, newBook)
	fmt.Println("Добавлена книга:", title, "автор:", author)
}

func (lib *Library) SearchByAuthor() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите имя автора для поиска:")
	author, err := reader.ReadString('\n')
	if err != nil {
        fmt.Println("Ошибка при чтении:", err)
        return
    }
	found := false
	author = strings.TrimSpace(author)
	for _, book := range lib.books {
		if book.Author == author {
			fmt.Println("--------------------------------------------------")
			fmt.Println("ID:", book.ID, "Название:", book.Title, "Автор:", book.Author, "Выдана:", book.IsIssued)
			fmt.Println("--------------------------------------------------")
			found = true
		}
	}

	if !found {
		fmt.Println("Книги данного автора не найдены.")
	}
}

func (lib *Library) SearchByTitle() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите название книги для поиска:")
	title, err := reader.ReadString('\n')
	if err != nil {
        fmt.Println("Ошибка при чтении:", err)
        return
    }
	found := false
	title = strings.TrimSpace(title)
	for _, book := range lib.books {
		if book.Title == title {
			fmt.Println("--------------------------------------------------")
			fmt.Println("ID:", book.ID, "Название:", book.Title, "Автор:", book.Author, "Выдана:", book.IsIssued)
			fmt.Println("--------------------------------------------------")
			found = true
		}
	}

	if !found {
		fmt.Println("Книги с таким названием не найдены.")
	}
}

func (lib *Library) IssueBook() {
	lib.PrintAllBooks()
	fmt.Println("Введите ID книги для выдачи:")
	var choice int
	fmt.Scanln(&choice)

	for i, book := range lib.books {
		if book.ID == choice {
			if lib.books[i].IsIssued {
				fmt.Println("Книга уже выдана.")
			} else {
				lib.books[i].IsIssued = true
				fmt.Println("Книга выдана:", book.Title)
			}
			return
		}
	}

	fmt.Println("Книга с таким ID не найдена.")
}

func (lib *Library) ReturnBook() {
	lib.PrintAllBooks()

	fmt.Println("Введите ID книги для возврата:")
	var choice int
	fmt.Scanln(&choice)

	for i, book := range lib.books {
		if book.ID == choice {
			if !lib.books[i].IsIssued {
				fmt.Println("Книга уже находится в библиотеке.")
			} else {
				lib.books[i].IsIssued = false
				fmt.Println("Книга возвращена:", book.Title)
			}
			return
		}
	}

	fmt.Println("Книга с таким ID не найдена.")
}

func (lib *Library) PrintAllBooks() {
	if len(lib.books) == 0 {
		fmt.Println("Библиотека пуста.")
		return
	}

	for _, book := range lib.books {
		fmt.Println("--------------------------------------------------")
		fmt.Println("ID:", book.ID, "Название:", book.Title, "Автор:", book.Author, "Выдана:", book.IsIssued)
		fmt.Println("--------------------------------------------------")
	}
}

func main() {
	library := Library{
		ID: 1,
	}
	library.AddBook("Мастер и Маргарита", "Михаил Булгаков")
	library.AddBook("Преступление и наказание", "Фёдор Достоевский")
	library.AddBook("Война и мир", "Лев Толстой")
	library.SearchByAuthor()
	library.SearchByTitle()
	library.IssueBook()
	library.ReturnBook()
}
