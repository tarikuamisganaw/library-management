package services

import (
	"errors"
	"fmt"

	"library_management/models"
)

// Contains business logic and data manipulation functions.

type Library struct {
	books   map[int]models.Book
	members map[int]models.Member
}

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

func NewLibrary() *Library {
	return &Library{
		books:   make(map[int]models.Book),
		members: make(map[int]models.Member),
	}
}

func (lib *Library) AddBook(book models.Book) {
	lib.books[book.ID] = book
}

func (lib *Library) RemoveBook(bookID int) {
	delete(lib.books, bookID)
}

func (lib *Library) BorrowBook(bookID int, memberID int) error {
	book, exists := lib.books[bookID]
	if !exists {
		return errors.New("book not found")
	}

	if book.Status == "Borrowed" {
		return errors.New("book is already borrowed")
	}

	member, exists := lib.members[memberID]
	if !exists {
		member = models.Member{ID: memberID, Name: fmt.Sprintf("Member %d", memberID)}
	}

	book.Status = "Borrowed"
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	lib.books[bookID] = book
	lib.members[memberID] = member
	return nil
}

func (lib *Library) ReturnBook(bookID int, memberID int) error {
	member, exists := lib.members[memberID]
	if !exists {
		return errors.New("member not found")
	}

	bookIndex := -1
	for i, book := range member.BorrowedBooks {
		if book.ID == bookID {
			bookIndex = i
			break
		}
	}

	if bookIndex == -1 {
		return errors.New("book not found in member's borrowed books")
	}

	book := member.BorrowedBooks[bookIndex]
	book.Status = "Available"
	member.BorrowedBooks = append(member.BorrowedBooks[:bookIndex], member.BorrowedBooks[bookIndex+1:]...)
	lib.books[bookID] = book
	lib.members[memberID] = member
	return nil
}

func (lib *Library) ListAvailableBooks() []models.Book {
	var availableBooks []models.Book
	for _, book := range lib.books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (lib *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, exists := lib.members[memberID]
	if !exists {
		return nil
	}
	return member.BorrowedBooks
}
