package models

// Book - REST GET model
type Book struct {
	ID       int64  `json:"id"`
	Isbn     string `json:"isbn"`
	Title    string `json:"title"`
	AuthorID int64  `json:"authorId"`
}

// Author - REST model
type Author struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
