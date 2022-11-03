package models

import (
	"time"
)

type TipoPlato string

const (
	Postre TipoPlato = "postre"
	Entrante 		 = "entrante"
	Primero 		 = "primero"
	Segundo 		 = "segundo"
)

type Plato struct {
	Nombre 			string
	Tipo 			TipoPlato
	Pasos 			[]PasoElaboracion
	TiempoInicio 	time.Time
}