package repo

import (
	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/model"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func GetNewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (br *BookRepository) AddBook(book *model.DBBook) {
	books := []model.DBBook{
		{Name: book.Name, Publisher: book.Publisher, Isbn: book.Isbn},
	}
	br.db.Create(&books)
}

func (br *BookRepository) UpdateBook(book *model.DBBook) {
	br.db.Model(&book).Where("isbn = ?", book.Isbn).Update("name", "publisher")

}

func (br *BookRepository) GetBook(isbn int) *model.DBBook {
	var book model.DBBook
	br.db.First(&book, isbn)
	return &book
}

func (br *BookRepository) GetAllBooks() ([]*model.DBBook, error) {
	books := make([]*model.DBBook, 0)
	err := br.db.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (br *BookRepository) RemoveBook(isbn int) {
	br.db.Delete(&model.DBBook{}, isbn)
}

func (br *BookRepository) AddReview(review *model.DBReview) {
	reviews := []model.DBReview{
		{Isbn: review.Isbn, Reviewer: review.Reviewer, Comment: review.Comment, Rating: review.Rating},
	}
	br.db.Create(&reviews)
}

func (br *BookRepository) GetAllReviews(isbn int) []model.DBReview {
	reviews := []model.DBReview{}
	br.db.Find(&reviews, isbn)

	return reviews
}
