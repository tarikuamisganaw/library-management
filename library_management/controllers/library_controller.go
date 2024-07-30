package controllers

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"library_management/models"
	"library_management/services"
)

//Handles console input and invokes the appropriate service methods.

type LibraryController struct {
	library *services.Library
}

func NewLibraryController() *LibraryController {
	return &LibraryController{
		library: services.NewLibrary(),
	}
}

func (lc *LibraryController) AddBook(reader *bufio.Reader) {
	fmt.Print("Enter book ID: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID. Please enter a number.")
		return
	}

	fmt.Print("Enter book title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Enter book author: ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)

	book := models.Book{
		ID:     id,
		Title:  title,
		Author: author,
		Status: "Available",
	}

	lc.library.AddBook(book)
	fmt.Println("Book added successfully.")
}

func (lc *LibraryController) RemoveBook(reader *bufio.Reader) {
	fmt.Print("Enter book ID to remove: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID. Please enter a number.")
		return
	}

	lc.library.RemoveBook(id)
	fmt.Println("Book removed successfully.")
}

func (lc *LibraryController) BorrowBook(reader *bufio.Reader) {
	fmt.Print("Enter book ID to borrow: ")
	bookIDStr, _ := reader.ReadString('\n')
	bookIDStr = strings.TrimSpace(bookIDStr)
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		fmt.Println("Invalid book ID. Please enter a number.")
		return
	}

	fmt.Print("Enter member ID: ")
	memberIDStr, _ := reader.ReadString('\n')
	memberIDStr = strings.TrimSpace(memberIDStr)
	memberID, err := strconv.Atoi(memberIDStr)
	if err != nil {
		fmt.Println("Invalid member ID. Please enter a number.")
		return
	}

	err = lc.library.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Book borrowed successfully.")
}

func (lc *LibraryController) ReturnBook(reader *bufio.Reader) {
	fmt.Print("Enter book ID to return: ")
	bookIDStr, _ := reader.ReadString('\n')
	bookIDStr = strings.TrimSpace(bookIDStr)
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		fmt.Println("Invalid book ID. Please enter a number.")
		return
	}

	fmt.Print("Enter member ID: ")
	memberIDStr, _ := reader.ReadString('\n')
	memberIDStr = strings.TrimSpace(memberIDStr)
	memberID, err := strconv.Atoi(memberIDStr)
	if err != nil {
		fmt.Println("Invalid member ID. Please enter a number.")
		return
	}

	err = lc.library.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Book returned successfully.")
}

func (lc *LibraryController) ListAvailableBooks() {
	books := lc.library.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("No available books.")
		return
	}

	fmt.Println("Available books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func (lc *LibraryController) ListBorrowedBooks(reader *bufio.Reader) {
	fmt.Print("Enter member ID: ")
	memberIDStr, _ := reader.ReadString('\n')
	memberIDStr = strings.TrimSpace(memberIDStr)
	memberID, err := strconv.Atoi(memberIDStr)
	if err != nil {
		fmt.Println("Invalid member ID. Please enter a number.")
		return
	}

	books := lc.library.ListBorrowedBooks(memberID)
	if len(books) == 0 {
		fmt.Println("No borrowed books for this member.")
		return
	}

	fmt.Println("Borrowed books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}
