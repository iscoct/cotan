package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPasosDeLasComandas(t *testing.T) {
	platos := createPlatos()
	pasos := createPasos()

	platos[0].Pasos = []PasoElaboracion{pasos[0], pasos[1]}
	platos[1].Pasos = []PasoElaboracion{pasos[2], pasos[3]}

	comanda_1 := Comanda{Platos: []Plato{platos[0], platos[0]}}
	comanda_2 := Comanda{Platos: []Plato{platos[1], platos[2]}}

	resultado_pasos := getPasosDeLasComandas([]Comanda{comanda_1, comanda_2})
	expected_pasos := []PasoElaboracion{pasos[0], pasos[1], pasos[0], pasos[1], pasos[2], pasos[3]}

	result := assert.ObjectsAreEqualValues(resultado_pasos, expected_pasos)
	assert.True(t, result, "Expected pasos y resultado pasos no son iguales")
}
