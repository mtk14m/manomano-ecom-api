package services

import (
	"testing"

	"github.com/mtk14m/manomano/internal/models"
)

type fakeProductRepository struct {
	createdProduct models.Product
	updatedProduct models.Product
	updatedID      int
}

func (f *fakeProductRepository) Create(p models.Product) (models.Product, error) {
	f.createdProduct = p
	return p, nil
}

func (f *fakeProductRepository) Update(id int, p models.Product) (models.Product, error) {
	f.updatedID = id
	f.updatedProduct = p
	return p, nil
}

func (f *fakeProductRepository) Delete(id int) error {
	return nil
}

func (f *fakeProductRepository) GetByID(id int) (models.Product, error) {
	return models.Product{}, nil
}

func (f *fakeProductRepository) GetAll() ([]models.Product, error) {
	return nil, nil
}

func TestCreateProduct_NormalizesProductBeforeSaving(t *testing.T) {
	fakeRepo := &fakeProductRepository{}
	service := NewProductService(fakeRepo)

	input := models.Product{
		Name:     " Perceuse visseuse ",
		Price:    99.99,
		Category: " Outillage ",
		InStock:  true,
	}

	_, err := service.CreateProduct(input)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if fakeRepo.createdProduct.Name != "Perceuse visseuse" {
		t.Errorf("expected Name to be %q, got %q", "Perceuse visseuse", fakeRepo.createdProduct.Name)
	}

	if fakeRepo.createdProduct.Category != "outillage" {
		t.Errorf("expected Category to be %q, got %q", "outillage", fakeRepo.createdProduct.Category)
	}

	if fakeRepo.createdProduct.Price != 99.99 {
		t.Errorf("expected Price to be %v, got %v", 99.99, fakeRepo.createdProduct.Price)
	}

	if fakeRepo.createdProduct.InStock != true {
		t.Errorf("expected InStock to be %v, got %v", true, fakeRepo.createdProduct.InStock)
	}
}

func TestUpdateProduct_NormalizesProductBeforeSaving(t *testing.T) {
	fakeRepo := &fakeProductRepository{}
	service := NewProductService(fakeRepo)

	id := 1
	input := models.Product{
		Name:     " Perceuse visseuse ",
		Price:    99.99,
		Category: " Outillage ",
		InStock:  true,
	}

	_, err := service.UpdateProduct(id, input)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if fakeRepo.updatedID != id {
		t.Errorf("expected Update to be called with id %d, got %d", id, fakeRepo.updatedID)
	}

	if fakeRepo.updatedProduct.Name != "Perceuse visseuse" {
		t.Errorf("expected Name to be %q, got %q", "Perceuse visseuse", fakeRepo.updatedProduct.Name)
	}

	if fakeRepo.updatedProduct.Category != "outillage" {
		t.Errorf("expected Category to be %q, got %q", "outillage", fakeRepo.updatedProduct.Category)
	}

	if fakeRepo.updatedProduct.Price != 99.99 {
		t.Errorf("expected Price to be %v, got %v", 99.99, fakeRepo.updatedProduct.Price)
	}

	if fakeRepo.updatedProduct.InStock != true {
		t.Errorf("expected InStock to be %v, got %v", true, fakeRepo.updatedProduct.InStock)
	}
}
