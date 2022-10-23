package models

import "time"

type Comanda struct {
	Mesa      Mesa
	Timestamp time.Time
	Platos    []Plato
	Bebidas   []Bebida
}
