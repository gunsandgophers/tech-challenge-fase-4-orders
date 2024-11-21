package repositories

import (
	"tech-challenge-fase-1/internal/infra/database"
	"tech-challenge-fase-1/internal/core/entities"
	valueobjects "tech-challenge-fase-1/internal/core/value_objects"
)

type CustomerRepositoryDB struct {
	conn database.ConnectionDB
}

func NewCustomerRepositoryDB(conn database.ConnectionDB) *CustomerRepositoryDB {
	return &CustomerRepositoryDB{conn: conn}
}

func (r *CustomerRepositoryDB) GetCustomerByCPF(cpf *valueobjects.CPF) (*entities.Customer, error) {
	sql := "SELECT id, name, email, cpf FROM customer WHERE cpf = $1"
	row := r.conn.QueryRow(sql, cpf.Value())
	return r.toEntity(row)
}

func (r *CustomerRepositoryDB) GetCustomerByEmail(email *valueobjects.Email) (*entities.Customer, error) {
	sql := "SELECT id, name, email, cpf FROM customer WHERE email = $1"
	row := r.conn.QueryRow(sql, email.Value())
	return r.toEntity(row)
}

func (r *CustomerRepositoryDB) Insert(customer *entities.Customer) error {
	sql := "INSERT INTO customer(id, name, email, cpf) VALUES($1, $2, $3, $4)"
	return r.conn.Exec(
		sql,
		customer.GetId(),
		customer.GetName(),
		customer.GetEmail().Value(),
		customer.GetCPF().Value(),
	)
}

func (r *CustomerRepositoryDB) toEntity(row database.RowDB) (*entities.Customer, error) {
	var id string
	var name string
	var email string
	var cpf string

	if err := row.Scan(&id, &name, &email, &cpf); err != nil {
		return nil, err
	}

	return entities.RestoreCustomer(id, name, email, cpf)
}

func (r *CustomerRepositoryDB) GetCustomerByID(id string) (*entities.Customer, error) {
	sql := "SELECT id, name, email, cpf FROM customer WHERE id = $1"
	row := r.conn.QueryRow(sql, id)

	customer, err := r.toEntity(row)

	if err != nil {
		if err.Error() == ErrNotFound {
			return nil, ErrCustomerNotFound
		}
		return nil, err
	}
	return customer, nil
}
