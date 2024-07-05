package dto

type BookDto struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BookResponse struct {
	Whole   []BookDto `json:"whole"`
	Counter int64     `json:"counter"`
}
