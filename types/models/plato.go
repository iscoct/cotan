package models

type TipoPlato int64

const (
	Postre TipoPlato = iota
	Entrante
	Primero
	Segundo
)

func (s TipoPlato) String() string {
	switch s {
	case Postre:
		return "postre"
	case Entrante:
		return "entrante"
	case Primero:
		return "primero"
	case Segundo:
		return "segundo"
	}
	return ""
}

type EstadoPlato int64

const (
	EnEspera EstadoPlato = iota
	Preparando
	ParaEntregar
	Entregado
)

func (s EstadoPlato) String() string {
	switch s {
	case EnEspera:
		return "en espera"
	case Preparando:
		return "preparándose"
	case ParaEntregar:
		return "para entregar"
	case Entregado:
		return "entregado"
	}
	return ""
}

type Plato struct {
	Nombre string
	Tipo   TipoPlato
	Estado EstadoPlato
}
