package model

type Tabler interface {
	TableName() string
}

type DBBook struct {
	Isbn      int    `json:"isbn"`
	Name      string `json:"name"`
	Publisher string `json:"publisher"`
}

func (DBBook) TableName() string {
	return "books"
}

type DBReview struct {
	Isbn     int    `json:"isbn"`
	Reviewer string `json:"reviewer"`
	Comment  string `json:"comment"`
	Rating   int    `json:"rating"`
}

func (DBReview) TableName() string {
	return "reviews"
}
