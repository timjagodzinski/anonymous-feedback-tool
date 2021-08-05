package main

import (
	"fmt"
	"github.com/timjagodzinski/go-crud-service/datatranspiler"
	"os"

	"github.com/rs/zerolog"
	"github.com/timjagodzinski/go-crud-service/persistencegateway"
)

func main() {
	l := zerolog.New(os.Stdout).With().Timestamp().Logger()

	c, err := persistencegateway.ConnectPGDatabase("root", "toor", "localhost:5432", "go_crud_example")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	source := persistencegateway.NewPGFruit(c, &l)
	transformer := datatranspiler.NewInventory(source, &l)

	s, err := transformer.AvailableFruits()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(s)
}
