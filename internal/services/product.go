package services

import (
	"strings"

	"github.com/mtk14m/manomano/internal/models"
)

type ProductRepository interface {
	Create(models.Product) (models.Product, error)
	Update(int, models.Product) (models.Product, error)
	Delete(int) error
	GetByID(int) (models.Product, error)
	GetAll() ([]models.Product, error)
}

type ProductService struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) CreateProduct(p models.Product) (models.Product, error) {
	p = normalizeProduct(p)
	return s.repo.Create(p)
}

func (s *ProductService) UpdateProduct(id int, p models.Product) (models.Product, error) {
	p = normalizeProduct(p)
	return s.repo.Update(id, p)
}

func (s *ProductService) DeleteProduct(id int) error {
	return s.repo.Delete(id)
}

func (s *ProductService) GetProductByID(id int) (models.Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.repo.GetAll()
}

func normalizeProduct(p models.Product) models.Product {
	p.Name = strings.TrimSpace(p.Name)
	p.Category = strings.ToLower(strings.TrimSpace(p.Category))
	return p
}
