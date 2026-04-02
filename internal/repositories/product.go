package repositories

import (
	"github.com/mtk14m/manomano/internal/models"
)

type ProductRepository struct {
	products []models.Product
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		products: productMockList,
	}
}

func (r *ProductRepository) GetAll() []models.Product {
	return r.products
}

func (r *ProductRepository) GetByID(id int) (models.Product, bool) {
	for _, product := range r.products {
		if product.ID == id {
			return product, true
		}
	}
	return models.Product{}, false
}

// Mock liste produits
var productMockList = []models.Product{
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
