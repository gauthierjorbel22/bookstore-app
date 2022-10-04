package controllers

type CreateBookInput struct {
	Title   string `json:"title" binding:"required"`
	Author  string `json:"author" binding:"required"`
	ISBN    string `json:"isbn" binding:"required"`
	Publisher string `json:"publisher" binding:"required"`
	Edition string `json:"edition" binding:"required"`
	Year string `json:"year" binding:"required`
	Picture string `json:"picture" binding:"required"`
}

type UpdateBookInput struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	ISBN    string `json:"isbn" binding:"required"`
	Publisher string `json:"publisher" binding:"required"`
	Edition string `json:"edition" binding:"required"`
	Year string `json:"year" binding:"required`
	Picture string `json:"picture" binding: "required"`
}
