package models

type TipoPlato string

const (
	Postre TipoPlato = "postre"
	Entrante 		 = "entrante"
	Primero 		 = "primero"
	Segundo 		 = "segundo"
)

type EstadoPlato string

const (
	EnEspera EstadoPlato = "en espera"
	Preparando 			 = "preparándose"
	ParaEntregar 		 = "para entregar"
	Entregado 			 = "entregado"
)

type Plato struct {
	Nombre string
	Tipo   TipoPlato
	Estado EstadoPlato
}
