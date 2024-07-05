package handler

//HTTP ile ilgili işlemleri yönetir.
import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"library2/internal/domain/dto"
	"library2/internal/domain/interf"
	"library2/internal/domain/model"
	"library2/pkg/validation"
	"net/http"
	"strconv"
)

type BookHandler struct {
	bookService interf.BookServiceInterface
	db          *gorm.DB
}

func NewBookHandler(bookService interf.BookServiceInterface) *BookHandler {
	return &BookHandler{
		bookService: bookService,
	}
}

func (h *BookHandler) GetDb() *gorm.DB {
	return h.db
}

func (h *BookHandler) GetAll(c *gin.Context) {
	var filters dto.BookDto
	if err := c.ShouldBindQuery(&filters); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bookResponse, err := h.bookService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bookResponse)

}

func (h *BookHandler) GetAllWithFilter(c *gin.Context) {
	var filters dto.BookFilter
	if err := c.ShouldBindQuery(&filters); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bookResponse, err := h.bookService.GetAllWithFilter(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bookResponse)
}
func (h *BookHandler) GetBook(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book, err := h.bookService.GetBook(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var input *model.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validation.Validate(*input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdBook, err := h.bookService.CreateBook(*input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, createdBook)
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	var input model.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedBook, err := h.bookService.UpdateBook(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedBook)
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.bookService.DeleteBook(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
