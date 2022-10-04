package models 

type Book struct{
	ID uint `json:"id" gorm:"primary_key"`
	Title string `json:"title"`
	Author string `json:"author"`
	ISBN string `json:"isbn"`
	Publisher string `json:"publisher"`
	Edition string `json:"edition"`
	Year string `json:"year"`
	Picture string `json:"picture"`
}
