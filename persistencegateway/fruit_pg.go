package persistencegateway

import (
	"context"
	"database/sql"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Fruits is a list of fruit.
type Fruits []Fruit

// Sometimes different domains might share a database and therefore should save the same connection pool.
// The connections pool should be protected against mischievous developers who want direct databse access.
// This makes testing hard and should therefore be prevented.
type pgFruit struct {
	conn *pgxpool.Pool
	log  *zerolog.Logger
}

// NewPGFruit instantiated a new postgres repository for the fruit table.
func NewPGFruit(conn *pgxpool.Pool, log *zerolog.Logger) FruitRepository {
	return &pgFruit{
		conn: conn,
		log:  log,
	}
}

// fruit represents the fruit table, the price is saved in cents.
type fruit struct {
	id     int
	name   string
	amount int
	price  int
}

// ID implements Fruit.ID
func (f *fruit) ID() int {
	return f.id
}

// Name implements Fruit.Name
func (f *fruit) Name() string {
	return f.name
}

// Amount implements Fruit.Amount
func (f *fruit) Amount() int {
	return f.amount
}

// Price implements Fruit.Price
func (f *fruit) Price() int {
	return f.price
}

// GetFruits Implements FruitRepository.GetFruits.
func (p *pgFruit) GetFruits(limit int) (fs []Fruit, err error) {
	q := "SELECT name, amount, price FROM fruit LIMIT $1"

	rows, err := p.conn.Query(context.Background(), q, limit)

	if err != nil {
		p.log.Err(err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		f := fruit{}
		err = rows.Scan(&f.name, &f.amount, &f.price)

		if err != nil {
			return nil, err
		}

		fs = append(fs, &f)
	}

	return fs, nil
}

// GetFruitByID Implements FruitRepository.GetFruitByID.
func (p *pgFruit) GetFruitByID(id int) (Fruit, error) {
	q := "SELECT name, amount, price FROM fruit WHERE ID = $1"
	f := fruit{}

	err := p.conn.QueryRow(context.Background(), q, id).Scan(&f.name, &f.amount, &f.price)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		p.log.Error().Err(err)
		return nil, err
	}

	return &f, nil
}
