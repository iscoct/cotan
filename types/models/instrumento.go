package models

import "time"

type Instrumento struct {
	Nombre string
}

type UsoInstrumento struct {
	Instrumento          Instrumento
	TiempoParaEstarLibre time.Time
}
