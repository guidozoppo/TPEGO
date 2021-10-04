package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectTipo(t *testing.T) {
	assert.Equal(t, detectTipo("TX"), "TX", "tipos no coinciden")
	assert.Equal(t, detectTipo("NN"), "NN", "tipos no coinciden")
	assert.Equal(t, detectTipo("Na"), "", "tipos no coinciden")
}

func TestDetectValue(t *testing.T) {
	assert.Equal(t, detectValue("AAAAAA"), "AA", "Values no coinciden")
}

func TestDetectLength(t *testing.T) {
	assert.Equal(t, detectLength("AA12"), 12, "Lengths no coinciden")
	assert.Equal(t, detectLength("AAaa"), 0, "Lengths no coinciden")
}

func TestIsInt(t *testing.T) {
	assert.Equal(t, isInt("asd"), false, "No es int")
	assert.Equal(t, isInt("12"), true, "Es int")
}
func TestIsLetter(t *testing.T) {
	assert.Equal(t, IsLetter("asd"), true, "Es letter")
	assert.Equal(t, IsLetter("12"), false, "No es letter")
}

func TestChain(t *testing.T) {
	var rsService ResultService
	rsService = &resultService{}
	r, _ := rsService.Chain("NN011")
	assert.Equal(t, r, "&{NN 1 1}", "Campos bien detectados")

	_, err := rsService.Chain("NNaa11")
	assert.Equal(t, err, "Cadena ingresada incoherente", "Error bien tirado")
}
