package repositories

import (
	"strings"
	"tech-challenge-fase-1/internal/core/entities"
	"tech-challenge-fase-1/internal/core/errors"
	"tech-challenge-fase-1/internal/infra/database"
)

type ProductRepositoryDB struct {
	conn database.ConnectionDB
}

func NewProductRepositoryDB(conn database.ConnectionDB) *ProductRepositoryDB {
	return &ProductRepositoryDB{conn: conn}
}

func (r *ProductRepositoryDB) Insert(product *entities.Product) error {
	sql := `
	INSERT INTO products(id, name, category, price, description, image)
	VALUES ($1, $2, $3, $4, $5, $6)
	`
	return r.conn.Exec(
		sql,
		product.GetId(),
		product.GetName(),
		product.GetCategory().String(),
		product.GetPrice(),
		product.GetDescription(),
		product.GetImage(),
	)
}

func (r *ProductRepositoryDB) FindProductByID(ID string) (*entities.Product, error) {
	sql := `
	SELECT
		id,
		name,
		category,
		price,
		description,
		image
	FROM
		products
	WHERE id = $1`
	row := r.conn.QueryRow(sql, ID)
	return r.toEntity(row)
}

func (r *ProductRepositoryDB) Update(product *entities.Product) error {
	sql := `
	UPDATE products
	SET
		name = $1,
		category = $2,
		price = $3,
		description = $4,
		image = $5,
		updated_at = NOW()
	WHERE id = $6
	`
	return r.conn.Exec(
		sql,
		product.GetName(),
		product.GetCategory().String(),
		product.GetPrice(),
		product.GetDescription(),
		product.GetImage(),
		product.GetId(),
	)
}

func (r *ProductRepositoryDB) Delete(ID string) error {
	query := "DELETE FROM products WHERE id = $1"
	err := r.conn.Exec(query, ID)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return errors.ErrProductNotFound
		}

		return err
	}

	return nil
}

func (r *ProductRepositoryDB) ListProducts() ([]*entities.Product, error) {
	sql := `
	SELECT
		id,
		name,
		category,
		price,
		description,
		image
	FROM
		products
	`
	rows, err := r.conn.Query(sql)
	if err != nil {
		return nil, err
	}
	var products []*entities.Product
	for rows.Next() {
		if p, err := r.toEntity(rows); err == nil {
			products = append(products, p)
		}
	}
	return products, nil
}

func (r *ProductRepositoryDB) FindProductByCategory(category entities.ProductCategory, page, size int) ([]*entities.Product, error) {
	sql := `
	SELECT
		id,
		name,
		category,
		price,
		description,
		image
	FROM
		products
	WHERE
		category LIKE $1
	LIMIT $2 OFFSET $3
	`
	rows, err := r.conn.Query(sql, category.String(), size, (page-1)*size)
	if err != nil {
		return nil, err
	}
	var products []*entities.Product
	for rows.Next() {
		if p, err := r.toEntity(rows); err == nil {
			products = append(products, p)
		}
	}
	return products, nil
}

func (r *ProductRepositoryDB) toEntity(row database.RowDB) (*entities.Product, error) {
	var id          string
	var name        string
	var category    entities.ProductCategory
	var price       float64
	var description string
	var image       string
	if err := row.Scan(&id, &name, &category, &price, &description, &image); err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, errors.ErrProductNotFound
		}
		return nil, err
	}

	return entities.RestoreProduct(id, name, category, price, description, image), nil
}
