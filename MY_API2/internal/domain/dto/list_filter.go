package dto

type BookFilter struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

/*type BookFilterResponse struct {
	Whole   []BookFilter `json:"whole"`
	Counter int64        `json:"counter"`
}*/
