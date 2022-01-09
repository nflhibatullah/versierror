package book

import (
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
	"layered/entities"
	"strconv"
)

type BookRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (br *BookRepository) Create(b entities.Book) (entities.Book, error) {

	if err := br.db.Save(&b).Error; err != nil {
		return b, err
	}

	return b, nil

}

func (br *BookRepository) GetAll() ([]entities.Book, error) {
	book := []entities.Book{}
	if err := br.db.Find(&book).Error; err != nil {
		log.Warn("Found database error:", err)
		return nil, err
	}
	return book, nil
}

func (br *BookRepository) Get(BookId int) ([]entities.Book, error) {

	book := []entities.Book{}

	if err := br.db.Find(&book, BookId).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (br *BookRepository) Delete(BookId int) (string, error) {
	var book = []entities.Book{}
	br.db.Delete(&book, BookId)
	return strconv.Itoa(BookId), nil
}

func (br *BookRepository) Update(newBook entities.Book, bookId int) (entities.Book, error) {
	var oldBook = entities.Book{}

	if err := br.db.Find(&oldBook, "?=id", bookId).Error; err != nil {
		return oldBook, err
	}

	oldBook.Title = newBook.Title

	return oldBook, nil
}
