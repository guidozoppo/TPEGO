package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

type Result struct {
	Type   string
	Length int
	Value  string
}

type ResultService interface {
	Chain(cadena string) (*Result, error)
}

type resultService struct {
}

func (rs *resultService) Chain(cadena string) (*Result, error) {
	tipo := detectTipo(cadena)
	value := detectValue(cadena)
	length := detectLength(cadena)
	lengthValue := len(value)

	if length > 0 {
		if ((tipo == "NN") && (isInt(value)) && (length == lengthValue)) ||
			((tipo == "TX") && (IsLetter(value)) && (length == lengthValue)) {
			return &Result{tipo, length, value}, nil
		}
	}
	return nil, errors.New("Cadena ingresada incoherente")
}

func detectTipo(tipo string) string {
	if tipo[:2] == "TX" {
		return "TX"
	} else if tipo[:2] == "NN" {
		return "NN"
	} else {
		return ""
	}
}

func detectValue(value string) string { //
	return value[4:]
}

func detectLength(length string) int {
	if sv, err := strconv.Atoi(length[2:4]); err == nil {
		return sv
	} else {
		return 0
	}
}

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func main() {
	/* var rsService services.ResultService
	rsService, err := services.NewResultService() */
	var rsService ResultService
	rsService = &resultService{}

	/* if err != nil {
		panic(err)
	} */

	r, err := rsService.Chain("NN0AA11")
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}
