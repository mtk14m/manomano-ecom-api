package repositories

import (
	"database/sql"

	"github.com/mtk14m/manomano/internal/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) GetAll() ([]models.Product, error) {
	rows, err := r.db.Query("SELECT id, name, price, category, in_stock FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Category, &p.InStock)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) GetByID(id int) (models.Product, error) {

	var p models.Product

	err := r.db.QueryRow(
		"SELECT id, name, price, category, in_stock FROM products WHERE id = $1",
		id,
	).Scan(&p.ID, &p.Name, &p.Price, &p.Category, &p.InStock)

	if err != nil {
		return models.Product{}, err
	}
	return p, nil
}

func (r *ProductRepository) Create(p models.Product) (models.Product, error) {

	var createdProduct models.Product
	err := r.db.QueryRow(
		`INSERT INTO products (name, price, category, in_stock)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id, name, price, category, in_stock`,
		p.Name,
		p.Price,
		p.Category,
		p.InStock,
	).Scan(
		&createdProduct.ID,
		&createdProduct.Name,
		&createdProduct.Price,
		&createdProduct.Category,
		&createdProduct.InStock,
	)

	if err != nil {
		return models.Product{}, err
	}

	return createdProduct, nil
}

func (r *ProductRepository) Update(id int, p models.Product) (models.Product, error) {
	var updatedProduct models.Product

	err := r.db.QueryRow(
		`
		UPDATE products
		SET name=$1, price=$2, category=$3, in_stock=$4
		WHERE id=$5
		RETURNING id, name, price, category, in_stock
		`,
		p.Name, p.Price, p.Category, p.InStock, id,
	).Scan(
		&updatedProduct.ID,
		&updatedProduct.Name,
		&updatedProduct.Price,
		&updatedProduct.Category,
		&updatedProduct.InStock,
	)

	if err != nil {
		return models.Product{}, err
	}

	return updatedProduct, nil
}

func (r *ProductRepository) Delete(id int) error {
	result, err := r.db.Exec(
		`DELETE FROM products WHERE id=$1`,
		id,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
