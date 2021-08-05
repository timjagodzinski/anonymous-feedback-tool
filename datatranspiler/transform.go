// Package datatranspiler provides services that changes the structure and content of a data source into a different representation.
//
// This is useful when the data coming out of a datebase/api is not directly  usable for a consuming service.
package datatranspiler

import (
	"encoding/json"
	"github.com/rs/zerolog"
	"github.com/timjagodzinski/go-crud-service/persistencegateway"
)

type Inventory interface {
	AvailableFruits() (string, error)
}

type inventory struct {
	repo persistencegateway.FruitRepository
	log  *zerolog.Logger
}

func NewInventory(repo persistencegateway.FruitRepository, log *zerolog.Logger) Inventory {
	return &inventory{
		repo: repo,
		log:  log,
	}
}

type fruit struct {
	Name   string  `json:"name"`
	Amount int     `json:"amount"`
	Price  float64 `json:"price"`
}

func (i *inventory) AvailableFruits() (string, error) {
	fs, err := i.repo.GetFruits(persistencegateway.LimitDefault)

	if err != nil {
		return "", err
	}

	var out []fruit

	for _, f := range fs {
		out = append(out, fruit{
			Name:   f.Name(),
			Amount: f.Amount(),
			Price:  centToEuro(f.Price()), // Is saved in cents, must be returned as euro
		})
	}

	b, err := json.Marshal(out)

	if err != nil {
		return "", err
	}

	return string(b), nil
}

func centToEuro(c int) (e float64) {
	e = float64(c)

	return e / 100
}
