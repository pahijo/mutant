package main

import (
	"encoding/json"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type MetaData interface{}

type DnaList struct {
	Dna []string
}

func (d *DnaList) AgregarLista(dl string) {
	d.Dna = append(d.Dna, dl)
}

func (d *DnaList) ToJson() ([]byte, error) {
	return json.Marshal(d)
}

type Dna struct {
	Dna [1]string
}

func (d *Dna) ToJson() ([]byte, error) {
	return json.Marshal(d)
}

// User schema of the user table
type Analisis struct {
	ID     int64 `json:"id"`
	Mutant bool  `json:"mutant"`
}

type XDna struct {
	Mutante int64 `json:"count_mutant_dna"`
	Humano  int64 `json:"count_human_dna"`
	RatioT  Ratio `json:"ratio"`
}

func (d *XDna) ToJson() ([]byte, error) {
	return json.Marshal(d)
}

type Ratio float64
