package containers

import (
	"cotan/types/models"
)

type ListaComandas []models.Comanda

func (lc ListaComandas) len() int {
	return len(lc)
}

func (lc ListaComandas) push(comanda models.Comanda) {
	for i, c:= range lc {
		if c.greater(comanda) {
			for j := len(lc); j > i; j-- {
				lc[j] := lc[j-1]
			}
			lc[i] = comanda
			break
		}
	}
}

func (lc ListaComandas) pop() models.Comanda {

	comanda := lc[0]

	for j := len(lc)-1; j > 0; j-- {
		lc[j-1] = lc[j]
	}

	return comanda
}