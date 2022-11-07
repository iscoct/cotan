package models

type Comanda struct {
	Platos []Plato
}

func getPasosDeLasComandas(comandas []Comanda) (pasos []PasoElaboracion) {
	for _, comanda := range comandas {
		for _, plato := range comanda.Platos {
			for _, paso := range plato.Pasos {
				pasos = append(pasos, paso)
			}
		}
	}

	return pasos
}
