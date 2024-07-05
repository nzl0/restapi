package db

import (
	"gorm.io/gorm"
	"library2/internal/domain/dto"
	"library2/internal/domain/interf"
	"library2/internal/domain/model"
)

type InMemoryBookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) interf.BookRepositoryInterface {
	return &InMemoryBookRepository{
		db: db,
	}
}
func (r *InMemoryBookRepository) GetBook(id int) (*model.Book, error) {
	var book model.Book
	if err := r.db.Where("id = ?", id).Find(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *InMemoryBookRepository) CreateBook(book model.Book) error {
	return r.db.Create(&book).Error
}

func (r *InMemoryBookRepository) UpdateBook(id int, book model.Book) error {
	if err := r.db.Where("id = ?", id).Updates(book).Error; err != nil {
		return err
	}
	return nil
}

func (r *InMemoryBookRepository) DeleteBook(id int) error {
	if err := r.db.Where("id = ?", id).Delete(&model.Book{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *InMemoryBookRepository) GetAll() ([]model.Book, int64, error) {
	var books []model.Book
	var ctr int64
	err := r.db.Model(&model.Book{}).Count(&ctr).Find(&books).Error
	if err != nil {
		return nil, ctr, err
	}
	return books, ctr, nil
}

func (r *InMemoryBookRepository) GetAllWithFilter(filters dto.BookFilter) ([]model.Book, int64, error) {
	var books []model.Book
	var ctr int64
	query := r.db.Model(&model.Book{})
	if filters.Title != "" {
		query.Where("title=?", filters.Title)
	}
	if filters.Author != "" {
		query.Where("author=?", filters.Author)
	}
	err := query.Count(&ctr).Find(&books).Error
	if err != nil {
		return nil, 0, err
	}
	return books, ctr, nil
}
