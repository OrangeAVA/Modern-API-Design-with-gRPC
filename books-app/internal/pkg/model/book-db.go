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
