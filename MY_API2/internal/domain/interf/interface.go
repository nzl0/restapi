package interf

import (
	"gorm.io/gorm"
	"library2/internal/domain/dto"
	"library2/internal/domain/model"
)

type BookRepositoryInterface interface { //veritabanı işlemleri
	GetAll() ([]model.Book, int64, error)
	GetBook(id int) (*model.Book, error)
	GetAllWithFilter(filters dto.BookFilter) ([]model.Book, int64, error)
	CreateBook(book model.Book) error
	UpdateBook(id int, book model.Book) error
	DeleteBook(id int) error

	//ControlUpdate(id int) (*model.Book, error)
}

type BookServiceInterface interface { //iş mantığı işlemleri
	GetAll() (*dto.BookResponse, error)
	GetBook(id int) (*model.Book, error)
	GetAllWithFilter(filters dto.BookFilter) (*dto.BookResponse, error)
	CreateBook(book model.Book) (*model.Book, error)
	UpdateBook(id int, book model.Book) (*model.Book, error)
	DeleteBook(id int) error
	CountDuplicate(book model.Book) error
}

type DatabaseInterface interface {
	Connection() (*gorm.DB, error)
	Close() error
	AutoMigrate(model ...interface{}) error
	Seed() error
}
