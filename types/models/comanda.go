package models

import (
	"time"
	"fmt"
)

type Comanda struct {
	Mesa      Mesa
	Timestamp time.Time
	Platos    []Plato
	Bebidas   []Bebida
}

func (c Comanda) status() string {
	status_entrantes := EnEspera
	entrantes_not_all_ready := false

	status_primeros := EnEspera
	primeros_not_all_ready := false
	
	status_segundos := EnEspera
	segundos_not_all_ready := false
	
	status_postres := EnEspera
	postres_not_all_ready := false

	for _, b:= range c.Platos {
		
		if b.Tipo == Entrante {
			switch status_entrantes {
			case EnEspera:
				status_entrantes = b.Estado
				if b.Estado != Entregado {
					entrantes_not_all_ready = true
				}
				break
			case Preparando:
				if (b.Estado == ParaEntregar || b.Estado == Entregado) && !entrantes_not_all_ready {
					status_entrantes = b.Estado
				}
				break
			case ParaEntregar:
				if b.Estado == Entregado && !entrantes_not_all_ready{
					status_entrantes = b.Estado
				}
				break
			case Entregado:
				if b.Estado != Entregado {
					entrantes_not_all_ready = true
					status_entrantes = b.Estado
				}
				break
			}
		} else if b.Tipo == Primero {
			switch status_primeros {
			case EnEspera:
				status_primeros = b.Estado
				if b.Estado != Entregado {
					primeros_not_all_ready = true
				}
				break
			case Preparando:
				if (b.Estado == ParaEntregar || b.Estado == Entregado) && !primeros_not_all_ready {
					status_primeros = b.Estado
				}
				break
			case ParaEntregar:
				if b.Estado == Entregado && !primeros_not_all_ready{
					status_primeros = b.Estado
				}
				break
			case Entregado:
				if b.Estado != Entregado {
					primeros_not_all_ready = true
					status_primeros = b.Estado
				}
				break
			}
		} else if b.Tipo == Segundo {
			switch status_segundos {
			case EnEspera:
				status_segundos = b.Estado
				if b.Estado != Entregado {
					segundos_not_all_ready = true
				}
				break
			case Preparando:
				if (b.Estado == ParaEntregar || b.Estado == Entregado) && !segundos_not_all_ready {
					status_segundos = b.Estado
				}
				break
			case ParaEntregar:
				if b.Estado == Entregado && !segundos_not_all_ready{
					status_segundos = b.Estado
				}
				break
			case Entregado:
				if b.Estado != Entregado {
					segundos_not_all_ready = true
					status_segundos = b.Estado
				}
				break
			}
		} else if b.Tipo == Postre {
			switch status_postres {
			case EnEspera:
				status_postres = b.Estado
				if b.Estado != Entregado {
					postres_not_all_ready = true
				}
				break
			case Preparando:
				if (b.Estado == ParaEntregar || b.Estado == Entregado) && !postres_not_all_ready {
					status_postres = b.Estado
				}
				break
			case ParaEntregar:
				if b.Estado == Entregado && !postres_not_all_ready{
					status_postres = b.Estado
				}
				break
			case Entregado:
				if b.Estado != Entregado {
					postres_not_all_ready = true
					status_postres = b.Estado
				}
				break
			}
		}
	}


	return fmt.Sprintf("Entrantes: %s\nPrimeros: %s\nSegundos: %s\nPostres: %s", status_entrantes, status_primeros, status_segundos, status_postres)
}


func (c Comanda) greater(comanda Comanda) bool {
	return c.Timestamp > comanda.Timestamp
}