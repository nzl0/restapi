package service

//Repository katmanını kullanarak veri işlemlerini gerçekleştirir
import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"library2/internal/domain/dto"
	"library2/internal/domain/interf"
	"library2/internal/domain/model"
)

/*var (
	ErrBookNotFound   = errors.New("book not found")
	ErrDuplicateTitle = errors.New("duplicate title")
)*/

type BookService struct {
	bookRepo interf.BookRepositoryInterface
	db       *gorm.DB
}

func NewBookService(repo interf.BookRepositoryInterface) *BookService {
	return &BookService{
		bookRepo: repo,
	}
}
func (u *BookService) GetDb() *gorm.DB {
	return u.db
}

func (u *BookService) GetAll() (*dto.BookResponse, error) {

	books, ctr, err := u.bookRepo.GetAll()
	if err != nil {
		return nil, err
	}
	var bookList []dto.BookDto
	for _, book := range books {
		oneBook := dto.BookDto{
			ID:     book.ID,
			Title:  book.Title,
			Author: book.Author,
		}
		bookList = append(bookList, oneBook)
	}
	return &dto.BookResponse{Whole: bookList, Counter: ctr}, nil
}

func (u *BookService) GetAllWithFilter(filters dto.BookFilter) (*dto.BookResponse, error) {
	bookList, ctr, err := u.bookRepo.GetAllWithFilter(filters)
	if err != nil {
		return nil, err
	}
	var bookListItem []dto.BookDto
	for _, books := range bookList {
		book := dto.BookDto{
			ID:     books.ID,
			Title:  books.Title,
			Author: books.Author,
		}
		bookListItem = append(bookListItem, book)
	}
	return &dto.BookResponse{Whole: bookListItem, Counter: ctr}, nil
}

func (u *BookService) GetBook(id int) (*model.Book, error) {
	book, err := u.bookRepo.GetBook(id)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (u *BookService) CountDuplicate(book model.Book) error { //for createBook
	var ctr int64 = 0
	if err := u.db.Model(&model.Book{}).Where("title = ? AND id != ?", book.Title, book.ID).Count(&ctr).Error; err != nil {
		return err
	}
	if ctr > 0 {
		return fmt.Errorf("%s is already exist", book.Title)
	}
	if err := u.db.Model(&model.Book{}).Where("id = ? AND title != ?", book.ID, book.Title).Count(&ctr).Error; err != nil {
		return err
	}
	if ctr > 0 {
		return fmt.Errorf("a book with ID %d but a different title already exists", book.ID)
	}
	return nil
}
func (u *BookService) CreateBook(book model.Book) (*model.Book, error) {
	err := u.CountDuplicate(book)
	if err != nil {
		return nil, err
	}
	if err := u.bookRepo.CreateBook(book); err != nil {
		return nil, err
	}
	return &book, nil
}

func (u *BookService) UpdateBook(id int, updatedBook model.Book) (*model.Book, error) {
	book, err := u.bookRepo.GetBook(id)
	if err != nil {
		return nil, err
	}

	book.Title = updatedBook.Title
	book.Author = updatedBook.Author

	err = u.bookRepo.UpdateBook(id, *book)
	if err != nil {
		return nil, err
	}
	return book, nil
}
func (u *BookService) DeleteBook(id int) error {
	book, err := u.bookRepo.GetBook(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("book with ID %d not found", id)
		}
		return err
	}

	return u.bookRepo.DeleteBook(book.ID)
}
