package models

import (
	"reflect"
	"testing"
	"time"
)

func createPlatos() []Plato {
	pizza := Plato{
		Nombre: "Pizza", TiempoInicio: time.Now(),
	}
	aneto := Plato{
		Nombre: "Aneto", TiempoInicio: time.Now().Local().Add(time.Hour),
	}
	arroz := Plato{
		Nombre: "Arroz", TiempoInicio: time.Now().Local().Add(-time.Hour),
	}

	return []Plato{pizza, aneto, arroz}
}

func createUsoInstrumentos() []UsoInstrumento {
	horno := UsoInstrumento{
		Instrumento: Instrumento{Nombre: "Horno"}, TiempoParaEstarLibre: time.Now(),
	}

	sarten := UsoInstrumento{
		Instrumento: Instrumento{Nombre: "Sartén"}, TiempoParaEstarLibre: horno.TiempoParaEstarLibre.Local().Add(-time.Hour),
	}

	microondas := UsoInstrumento{
		Instrumento: Instrumento{Nombre: "Microondas"}, TiempoParaEstarLibre: horno.TiempoParaEstarLibre.Local().Add(time.Hour),
	}

	return []UsoInstrumento{horno, sarten, microondas}
}

func createPasos() []PasoElaboracion {
	amasar_pizza := PasoElaboracion{Nombre: "Amasar"}
	hornear_pizza := PasoElaboracion{Nombre: "Hornear"}
	empanar_aneto := PasoElaboracion{Nombre: "Empanar"}
	freir_aneto := PasoElaboracion{Nombre: "Freir"}

	return []PasoElaboracion{amasar_pizza, hornear_pizza, empanar_aneto, freir_aneto}
}

func TestAddTiempoDePasoAInstrumento(t *testing.T) {
	current_moment := time.Now()
	duracion_paso := 2 * time.Minute
	instrumento := Instrumento{Nombre: "Horno"}
	uso_instrumento := UsoInstrumento{Instrumento: instrumento, TiempoParaEstarLibre: current_moment}
	paso := PasoElaboracion{Duracion: duracion_paso, TiempoInicio: current_moment}

	addTiempoDePasoAInstrumento(&uso_instrumento, paso)
	diferencia_en_segundos := uso_instrumento.TiempoParaEstarLibre.Sub(current_moment).Seconds()

	if diferencia_en_segundos != duracion_paso.Seconds() {
		t.Fatalf("TestAddTiempoDePasoAInstrumento failed")
	}
}

func TestRemovePasoDePasosNoOptimizadosPrimerElemento(t *testing.T) {
	pos_to_remove := 0
	pasos := createPasos()

	expected_pasos := append([]PasoElaboracion{}, pasos[1:]...)
	pasos = removePasoDePasosNoOptimizados(pasos, pos_to_remove)

	if !reflect.DeepEqual(expected_pasos, pasos) {
		t.Fatalf("TestRemovePasoDePasosNoOptimizadosPrimerElemento failed")
	}
}

func TestRemovePasoDePasosNoOptimizadosUltimoElemento(t *testing.T) {
	pos_to_remove := 0
	pasos := createPasos()

	expected_pasos := append([]PasoElaboracion{}, pasos[1:]...)
	pasos = removePasoDePasosNoOptimizados(pasos, pos_to_remove)

	if !reflect.DeepEqual(expected_pasos, pasos) {
		t.Fatalf("TestRemovePasoDePasosNoOptimizadosUltimoElemento failed")
	}
}

func TestGetIndexDePasosSinDependencias(t *testing.T) {
	pasos := createPasos()

	pasos[1].DependenciaPasos = []int{pasos[0].ID}
	pasos[3].DependenciaPasos = []int{pasos[2].ID}

	indices := getIndexDePasosSinDependencias(pasos)
	expected_indices := []int{0, 2}

	if !reflect.DeepEqual(indices, expected_indices) {
		t.Fatalf("TestGetIndexDePasosSinDependencias failed")
	}
}

func TestGetPasosSinDependenciasOrdenadas(t *testing.T) {
	pasos := createPasos()
	pasos[1].DependenciaPasos = []int{pasos[0].ID}
	pasos[3].DependenciaPasos = []int{pasos[2].ID}
	indices_pasos_sin_deps := []int{0, 2}
	pasos[0].TiempoInicio = pasos[0].TiempoInicio.Add(2 * time.Hour)

	expected_indices := make([]int, len(indices_pasos_sin_deps))
	copy(expected_indices, indices_pasos_sin_deps)
	expected_indices[0], expected_indices[1] = expected_indices[1], expected_indices[0]

	getPasosSinDependenciasOrdenados(pasos, indices_pasos_sin_deps)

	if !reflect.DeepEqual(expected_indices, indices_pasos_sin_deps) {
		t.Fatalf("TestGetPasosSinDependenciasOrdenadas failed")
	}
}

func TestGetPosSiguientePasoEInstrumento(t *testing.T) {
	pasos := createPasos()
	instrumentos := createUsoInstrumentos()
	tabla := UsoInstrumento{
		Instrumento: Instrumento{Nombre: "Tabla"}, TiempoParaEstarLibre: time.Now(),
	}
	instrumentos = append(instrumentos, tabla)
	pasos[1].DependenciaPasos = []int{pasos[0].ID}
	pasos[3].DependenciaPasos = []int{pasos[2].ID}
	pasos[0].Instrumento = tabla.Instrumento
	pasos[2].Instrumento = tabla.Instrumento
	pasos[1].Instrumento = instrumentos[0].Instrumento
	pasos[3].Instrumento = instrumentos[1].Instrumento
	indices_pasos_sin_deps := []int{0, 2}
	pasos[0].TiempoInicio = pasos[0].TiempoInicio.Add(1 * time.Hour)

	pos_instrumento, pos_paso := GetPosSiguientePasoEInstrumento(pasos, indices_pasos_sin_deps, instrumentos)

	expected_pos_instrumento, expected_pos_paso := 3, 2

	if expected_pos_instrumento != pos_instrumento {
		t.Errorf("TestGetPosSiguientePasoEInstrumento pos_instrumento no es el que se esperaba")
	}

	if expected_pos_paso != pos_paso {
		t.Errorf("TestGetPosSiguientePasoEInstrumento pos_paso no es el que se esperaba")
	}
}

func TestGetPosSiguientePasoEInstrumento2(t *testing.T) {
	pasos := createPasos()
	instrumentos := createUsoInstrumentos()
	tabla := UsoInstrumento{
		Instrumento: Instrumento{Nombre: "Tabla"}, TiempoParaEstarLibre: time.Now(),
	}
	instrumentos = append(instrumentos, tabla)
	pasos[0].Instrumento = tabla.Instrumento
	pasos[2].Instrumento = tabla.Instrumento
	pasos[1].Instrumento = instrumentos[0].Instrumento
	pasos[3].Instrumento = instrumentos[1].Instrumento
	indices_pasos_sin_deps := []int{0, 1, 2, 3}
	pasos[0].TiempoInicio = pasos[0].TiempoInicio.Add(1 * time.Hour)
	instrumentos[0].TiempoParaEstarLibre = instrumentos[0].TiempoParaEstarLibre.Add(2 * time.Hour)
	instrumentos[3].TiempoParaEstarLibre = instrumentos[0].TiempoParaEstarLibre.Add(30 * time.Minute)

	pos_instrumento, pos_paso := GetPosSiguientePasoEInstrumento(pasos, indices_pasos_sin_deps, instrumentos)

	expected_pos_instrumento, expected_pos_paso := 1, 3

	if expected_pos_instrumento != pos_instrumento {
		t.Errorf("TestGetPosSiguientePasoEInstrumento2 pos_instrumento no es el que se esperaba")
	}

	if expected_pos_paso != pos_paso {
		t.Errorf("TestGetPosSiguientePasoEInstrumento2 pos_paso no es el que se esperaba")
	}
}

func TestFindPasoParaInstrumento(t *testing.T) {
	pasos := createPasos()
	instrumentos := createUsoInstrumentos()
	indices := []int{0, 1, 2}
	pasos[0].DependenciaPasos = []int{1}
	pasos[1].DependenciaPasos = []int{2}
	pasos[2].DependenciaPasos = []int{0}
	pasos[0].Instrumento = instrumentos[0].Instrumento
	pasos[1].Instrumento = instrumentos[0].Instrumento
	pasos[2].Instrumento = instrumentos[0].Instrumento

	pos_paso := findPasoParaInstrumento(pasos, indices, instrumentos[0].Instrumento)

	if pos_paso != -1 {
		t.Fatalf("TestFindPasoParaInstrumento no devuelve la posición del paso esperada")
	}
}

func TestFindPasoParaInstrumento2(t *testing.T) {
	pasos := createPasos()
	instrumentos := createUsoInstrumentos()
	indices := []int{0, 1, 2}

	pos_paso := findPasoParaInstrumento(pasos, indices, instrumentos[0].Instrumento)

	if pos_paso != -1 {
		t.Fatalf("TestFindPasoParaInstrumento2 no devuelve la posición del paso esperada")
	}
}

func TestFindPasoParaInstrumento3(t *testing.T) {
	pasos := createPasos()
	instrumentos := createUsoInstrumentos()
	indices := []int{0, 1, 2}
	pasos[0].Instrumento = instrumentos[0].Instrumento
	pasos[1].Instrumento = instrumentos[0].Instrumento
	pasos[2].Instrumento = instrumentos[0].Instrumento

	pos_paso := findPasoParaInstrumento(pasos, indices, instrumentos[0].Instrumento)

	if pos_paso == -1 {
		t.Fatalf("TestFindPasoParaInstrumento3 no devuelve la posición del paso esperada")
	}
}

func TestGetPasosOptimizados(t *testing.T) {
	current_moment := time.Now()
	tabla := Instrumento{Nombre: "Tabla"}
	uso_tabla := UsoInstrumento{Instrumento: tabla, TiempoParaEstarLibre: current_moment.Add(1 * time.Hour)}
	tabla_2 := Instrumento{Nombre: "Tabla 2"}
	uso_tabla_2 := UsoInstrumento{Instrumento: tabla_2, TiempoParaEstarLibre: current_moment.Add(1*time.Hour + 10*time.Minute)}
	horno := Instrumento{Nombre: "Horno"}
	uso_horno := UsoInstrumento{Instrumento: horno, TiempoParaEstarLibre: current_moment}
	sarten := Instrumento{Nombre: "Sartén"}
	uso_sarten := UsoInstrumento{Instrumento: sarten, TiempoParaEstarLibre: current_moment}
	amasar_pizza_1 := PasoElaboracion{Nombre: "Amasar", Instrumento: tabla, Duracion: 2 * time.Minute, ID: 1}
	amasar_pizza_2 := PasoElaboracion{Nombre: "Amasar", Instrumento: tabla, Duracion: 2 * time.Minute, ID: 2}
	hornear_pizza_1 := PasoElaboracion{Nombre: "Hornear", Instrumento: horno, Duracion: 30 * time.Minute, DependenciaPasos: []int{amasar_pizza_1.ID}}
	hornear_pizza_2 := PasoElaboracion{Nombre: "Hornear", Instrumento: horno, Duracion: 30 * time.Minute, DependenciaPasos: []int{amasar_pizza_2.ID}}
	empanar := PasoElaboracion{Nombre: "Empanar", Instrumento: tabla_2, Duracion: 3 * time.Minute, ID: 3}
	freir := PasoElaboracion{Nombre: "Freir", Instrumento: sarten, Duracion: 10 * time.Minute, DependenciaPasos: []int{empanar.ID}}

	pizza_1 := Plato{Nombre: "Pizza", Pasos: []PasoElaboracion{amasar_pizza_1, hornear_pizza_1}, TiempoInicio: current_moment}
	pizza_2 := Plato{Nombre: "Pizza", Pasos: []PasoElaboracion{amasar_pizza_2, hornear_pizza_2}, TiempoInicio: current_moment}
	aneto := Plato{Nombre: "Aneto", Pasos: []PasoElaboracion{empanar, freir}, TiempoInicio: current_moment.Add(30 * time.Second)}
	comanda_1 := Comanda{Platos: []Plato{pizza_1, pizza_2}}
	comanda_2 := Comanda{Platos: []Plato{aneto}}
	lista_comandas := []Comanda{comanda_1, comanda_2}

	pasos_optimizados := getPasosOptimizados([]UsoInstrumento{uso_tabla, uso_sarten, uso_tabla_2, uso_horno}, lista_comandas)
	expected_pasos_optimizados := []PasoElaboracion{
		amasar_pizza_2, amasar_pizza_1, hornear_pizza_1, empanar, freir, hornear_pizza_2,
	}

	for i, paso := range pasos_optimizados {
		if paso.ID != expected_pasos_optimizados[i].ID {
			t.Fatalf("TestGetPasosOptimizados failed")
		}
	}
}
