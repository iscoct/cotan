package models

import (
	"time"
)

type PasoElaboracion struct {
	ID               int
	Nombre           string
	Instrumento      Instrumento
	Duracion         time.Duration
	TiempoInicio     time.Time
	TiempoFinal      time.Time
	DependenciaPasos []int
}
