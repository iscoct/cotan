package models

import (
	"time"
)

type PasoElaboracion struct {
	Nombre 			string
	Instrumento 	Instrumento
	Duracion		int
	TiempoInicio	time.Time
	TiempoFinal		time.Time
}