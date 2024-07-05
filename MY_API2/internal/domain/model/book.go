package model

type Book struct {
	ID     int    `json:"id" validate:"required,gte=1,lte=100" gorm:"primary_key;auto_increment"`
	Title  string `json:"title" validate:"required,max=100" gorm:"type:varchar(100)"`
	Author string `json:"author" validate:"required,max=200"  gorm:"type:varchar(200)"`
}
