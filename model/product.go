package model

type Product struct {
	Id          int64   `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name        string  `json:"name" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Quantity    int64   `json:"quantity"`
}

type ProductResponse struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Quantity    int64   `json:"quantity"`
}
