package models

import "time"

type Comanda struct {
	Mesa      Mesa
	Timestamp time.Time
	Platos    []Plato
	Bebidas   []Bebida
}

func (c Comanda) greater(comanda Comanda) bool {
	return c.Timestamp > comanda.Timestamp
}