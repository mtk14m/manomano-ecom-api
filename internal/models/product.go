package models

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
	InStock  bool    `json:"in_stock"`
}


//Mock liste produits
var ProductMockList []Product = []Product{
		{
			ID:       1,
			Name:     "Perceuse visseuse",
			Price:    89.99,
			Category: "outillage",
			InStock:  true,
		},
		{
			ID:       2,
			Name:     "Aspirateur robot",
			Price:    249.99,
			Category: "électroménager",
			InStock:  false,
		},
		{
			ID:       3,
			Name:     "Casque audio sans fil",
			Price:    129.99,
			Category: "high-tech",
			InStock:  true,
		},
}
