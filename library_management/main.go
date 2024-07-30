package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"library_management/controllers"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	controller := controllers.NewLibraryController()

	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add a new book")
		fmt.Println("2. Remove an existing book")
		fmt.Println("3. Borrow a book")
		fmt.Println("4. Return a book")
		fmt.Println("5. List all available books")
		fmt.Println("6. List all borrowed books by a member")
		fmt.Println("7. Exit")
		fmt.Print("Enter your choice: ")

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid choice. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			controller.AddBook(reader)
		case 2:
			controller.RemoveBook(reader)
		case 3:
			controller.BorrowBook(reader)
		case 4:
			controller.ReturnBook(reader)
		case 5:
			controller.ListAvailableBooks()
		case 6:
			controller.ListBorrowedBooks(reader)
		case 7:
			fmt.Println("Exiting the system.")
			return
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 7.")
		}
	}
}
