package model

type Book struct {
	Isbn      int    `json:"isbn"`
	Name      string `json:"name"`
	Publisher string `json:"publisher"`
}

type Review struct {
	Isbn     int    `json:"isbn"`
	Reviewer string `json:"reviewer"`
	Comment  string `json:"comment"`
	Rating   int    `json:"rating"`
}
