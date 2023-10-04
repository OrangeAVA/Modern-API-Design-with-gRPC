package model

type Book struct {
	Isbn      int    `json:"isbn"`
	Name      string `json:"name"`
	Publisher string `json:"publisher"`
}
