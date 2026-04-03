package dtos

type CreateProductDto struct {
	Name     string  `json:"name" binding:"required,min=1"`
	Price    float64 `json:"price" binding:"required,gt=0"`
	Category string  `json:"category" binding:"required,min=1"`
	InStock  bool    `json:"in_stock"`
}
