// Package persistencegateway contains sources of data, relational databases, document stores or 3rd party APIs etc.
//
// These sources are exclusively communicated with, with the provided Interfaces. No direct access can be cranted to the outside via queries.
// The goal is to protect these datasources from random, unrestircted access.
package persistencegateway

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	LimitDefault    = 30
	OrderAscending  = "ASC"
	OrderDescending = "DESC"
	OrderDefault    = OrderAscending
)

func ConnectPGDatabase(user string, password string, host string, database string) (*pgxpool.Pool, error) {

	conn, err := pgxpool.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s/%s", user, password, host, database))

	if err != nil {
		return nil, err
	}

	return conn, nil
}

// FruitRepository provides access to information about fruit we have.
//
// GetFruits gives the list of fruits.
// GetFruitsByID gets a fruit by its ID.
type FruitRepository interface {
	GetFruits(limit int) (fs []Fruit, err error)
	GetFruitByID(id int) (f Fruit, err error)
}

// Fruit provides access to fruit data.
type Fruit interface {
	ID() int
	Name() string
	Amount() int
	Price() int
}
