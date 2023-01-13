package models

import (
	"cotan/cmd/pkg/logger"
	"time"
)

func addTiempoDePasoAInstrumento(instrumento *UsoInstrumento, paso PasoElaboracion) {
	instrumento.TiempoParaEstarLibre = paso.TiempoInicio.Add(paso.Duracion)
}

func indexDependenciaPaso(paso PasoElaboracion, id_paso int) int {
	pos := -1
	i := 0

	for i < len(paso.DependenciaPasos) && pos == -1 {
		if paso.DependenciaPasos[i] == id_paso {
			pos = i
		} else {
			i++
		}
	}

	logger.Standard.Info("Pos devuelta: ", pos)
	logger.Standard.Info("Id paso: ", id_paso)
	logger.Standard.Info("Paso elaboracion: ", paso)

	return pos
}

func removePasoDePasosNoOptimizados(pasos []PasoElaboracion, pos int) []PasoElaboracion {
	id_paso := pasos[pos].ID

	for i, paso := range pasos {
		pos_dependencia := indexDependenciaPaso(paso, id_paso)

		if pos_dependencia != -1 {
			if pasos[i].TiempoInicio.Before(pasos[pos].TiempoInicio) {
				pasos[i].TiempoInicio = pasos[pos].TiempoInicio
			}

			pasos[i].TiempoInicio = pasos[i].TiempoInicio.Add(pasos[pos].Duracion)
			pasos[i].DependenciaPasos = append(paso.DependenciaPasos[:pos_dependencia], paso.DependenciaPasos[pos_dependencia+1:]...)
		}
	}

	return append(pasos[:pos], pasos[pos+1:]...)
}

func getIndexDePasosSinDependencias(pasos []PasoElaboracion) (indices []int) {
	for i, val := range pasos {
		if len(val.DependenciaPasos) == 0 {
			indices = append(indices, i)
		}
	}

	return indices
}

func getPasosSinDependenciasOrdenados(pasos []PasoElaboracion, indices_de_pasos_sin_dep []int) {
	var isDone = false

	for !isDone {
		isDone = true
		var i = 0

		for i < len(indices_de_pasos_sin_dep)-1 {
			paso_actual := pasos[indices_de_pasos_sin_dep[i]]
			paso_siguiente := pasos[indices_de_pasos_sin_dep[i+1]]

			if paso_actual.TiempoInicio.After(paso_siguiente.TiempoInicio) {
				indices_de_pasos_sin_dep[i], indices_de_pasos_sin_dep[i+1] = indices_de_pasos_sin_dep[i+1], indices_de_pasos_sin_dep[i]

				isDone = false
			}

			i++
		}
	}
}

func findPasoParaInstrumento(pasos []PasoElaboracion, indices []int, instrumento Instrumento) int {
	pos_encontrado := -1
	i := 0

	for i < len(indices) {
		if pasos[indices[i]].Instrumento.Nombre == instrumento.Nombre && len(pasos[indices[i]].DependenciaPasos) == 0 {
			pos_encontrado = indices[i]
		}

		i++
	}

	return pos_encontrado
}

func GetPosSiguientePasoEInstrumento(pasos []PasoElaboracion, indices []int, instrumentos []UsoInstrumento) (pos_instrumento int, pos_paso int) {
	var max_tiempo time.Time
	max_menor_tiempo := time.Now().AddDate(10, 0, 0)

	for i, instrumento := range instrumentos {
		pos_paso_para_instrumento := findPasoParaInstrumento(pasos, indices, instrumento.Instrumento)

		if pos_paso_para_instrumento != -1 {
			if pasos[pos_paso_para_instrumento].TiempoInicio.After(instrumento.TiempoParaEstarLibre) {
				max_tiempo = pasos[pos_paso_para_instrumento].TiempoInicio
			} else {
				max_tiempo = instrumento.TiempoParaEstarLibre
			}

			if max_tiempo.Before(max_menor_tiempo) {
				pos_instrumento = i
				pos_paso = pos_paso_para_instrumento
				max_menor_tiempo = max_tiempo
			}
		}
	}

	return pos_instrumento, pos_paso
}

func addAPasosElTiempoParaEstarLibreDeSuInstrumento(instrumentos []UsoInstrumento, pasos []PasoElaboracion) {
	for i, paso := range pasos {
		encontrado := false
		j := 0

		for len(instrumentos) > j && !encontrado {
			if paso.Instrumento.Nombre == instrumentos[j].Instrumento.Nombre {
				encontrado = true
			} else {
				j++
			}
		}

		if encontrado && pasos[i].TiempoInicio.Before(instrumentos[j].TiempoParaEstarLibre) {
			pasos[i].TiempoInicio = instrumentos[j].TiempoParaEstarLibre
		}
	}
}

func printInitialInfo(comandas []Comanda, instrumentos []UsoInstrumento) {
	logger.Standard.Info("Comandas: ", comandas)
	logger.Standard.Info("Instrumentos: ", instrumentos)
}

func printPasosInfo(pasos_no_optimizados []PasoElaboracion, indices_pasos_sin_deps []int) {
	logger.Standard.Info("Pasos no optimizados: ", pasos_no_optimizados)
	logger.Standard.Info("Índices de pasos sin dependencias antes de ordenarlos: ", indices_pasos_sin_deps)
}

func getPasosOptimizados(instrumentos []UsoInstrumento, comandas []Comanda) (pasos_optimizados []PasoElaboracion) {
	pasos_no_optimizados := getPasosDeLasComandas(comandas)
	indices_pasos_sin_deps := getIndexDePasosSinDependencias(pasos_no_optimizados)

	printInitialInfo(comandas, instrumentos)
	printPasosInfo(pasos_no_optimizados, indices_pasos_sin_deps)

	addAPasosElTiempoParaEstarLibreDeSuInstrumento(instrumentos, pasos_no_optimizados)
	getPasosSinDependenciasOrdenados(pasos_no_optimizados, indices_pasos_sin_deps)

	printPasosInfo(pasos_no_optimizados, indices_pasos_sin_deps)

	for len(pasos_no_optimizados) > 0 {
		pos_instrumento, pos_paso := GetPosSiguientePasoEInstrumento(pasos_no_optimizados, indices_pasos_sin_deps, instrumentos)

		addTiempoDePasoAInstrumento(&instrumentos[pos_instrumento], pasos_no_optimizados[pos_paso])
		pasos_optimizados = append(pasos_optimizados, pasos_no_optimizados[pos_paso])
		pasos_no_optimizados = removePasoDePasosNoOptimizados(pasos_no_optimizados, pos_paso)
		indices_pasos_sin_deps = getIndexDePasosSinDependencias(pasos_no_optimizados)

		printPasosInfo(pasos_no_optimizados, indices_pasos_sin_deps)

		addAPasosElTiempoParaEstarLibreDeSuInstrumento(instrumentos, pasos_no_optimizados)
		getPasosSinDependenciasOrdenados(pasos_no_optimizados, indices_pasos_sin_deps)

		printPasosInfo(pasos_no_optimizados, indices_pasos_sin_deps)
		logger.Standard.Info("Pasos optimizados: ", pasos_optimizados)
	}

	return pasos_optimizados
}
