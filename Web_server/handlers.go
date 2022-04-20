package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

func isMutant(w http.ResponseWriter, r *http.Request) {
	var mutant DnaList
	isValido := true
	const size = 6
	var matrizDna = [size][size]string{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Se presento un error")
		isValido = false
		return
	}
	json.Unmarshal(reqBody, &mutant)

	//La información ingresada sea valida
	matrizDna, isValido, message := isValid(mutant, 6)
	if !isValido {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, message)
		return
	}

	isMutante := 0
	c := make(chan int, 3)
	var wg sync.WaitGroup
	c <- 1
	wg.Add(1)
	go func() {
		isMutante = isMutante + isMutantH(matrizDna, &wg, c)
	}()
	c <- 1
	wg.Add(1)
	go func() {
		isMutante = isMutante + isMutantV(matrizDna, &wg, c)
	}()
	c <- 1
	wg.Add(1)
	go func() {
		isMutante = isMutante + isMutantO(matrizDna, &wg, c)
	}()

	wg.Wait()

	if isMutante >= 2 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Es mutante")
		Create(w, true)
	} else {
		w.WriteHeader(http.StatusForbidden)
		Create(w, false)
		fmt.Fprintf(w, "No es mutante")
	}
}

func isValid(mutant DnaList, tam int) ([6][6]string, bool, string) {
	var matrizDna = [6][6]string{}
	for index, tarea := range mutant.Dna {
		//Valida el tamaño de cada registro corresponda al total de items del arreglo NXN
		if len(tarea) != len(mutant.Dna) {
			//w.WriteHeader(http.StatusBadRequest)
			//fmt.Fprintf(w, "La información ingresada debe corresponder a una matrix NxN")
			return matrizDna, false, "La información ingresada debe corresponder a una matrix NxN"
		}
		//Valida que solo tenga las letras permitidas  A T C G
		valorPermitido := true
		for indexInterno, r := range tarea {
			if !(strings.Contains("A", string(r)) || strings.Contains("T", string(r)) || strings.Contains("C", string(r)) || strings.Contains("G", string(r))) {
				valorPermitido = false
			} else {
				matrizDna[index][indexInterno] = string(r)
			}
		}
		if !valorPermitido {
			//w.WriteHeader(http.StatusBadRequest)
			//fmt.Fprintf(w, "Valores no permitidos en la posición %v\n", index)
			return matrizDna, false, "Valores no permitidos en la posición %v\n"
		}
	}
	return matrizDna, true, ""
}

func isMutantH(matrizDna [6][6]string, wg *sync.WaitGroup, c chan int) int {
	defer wg.Done()
	isMutante := 0
	//Validar Mutante Horizontal
	for f := 0; f < len(matrizDna); f++ {
		isConcidente := 0
		for c := 0; c < len(matrizDna)-1; c++ {
			if matrizDna[f][c] == matrizDna[f][c+1] {
				isConcidente++
				if isConcidente == 3 {
					isMutante++
				}
			} else {
				isConcidente = 0
			}
		}
	}
	<-c
	return isMutante
}

func isMutantV(matrizDna [6][6]string, wg *sync.WaitGroup, c chan int) int {
	defer wg.Done()
	isMutante := 0
	//Validar Mutante Vertical
	for f := 0; f < len(matrizDna); f++ {
		isConcidente := 0
		for c := 0; c < len(matrizDna)-1; c++ {
			if matrizDna[c][f] == matrizDna[c+1][f] {
				isConcidente++
				if isConcidente == 3 {
					isMutante++
				}
			} else {
				isConcidente = 0
			}
		}
	}
	<-c
	return isMutante
}

func isMutantO(matrizDna [6][6]string, wg *sync.WaitGroup, c chan int) int {
	defer wg.Done()
	//Validar Mutante Oblicua
	isMutante := 0
	isConcidente := 0
	for f := 0; f < len(matrizDna); f += 2 {
		if matrizDna[f][f] == matrizDna[f+1][f+1] {
			isConcidente++
			if isConcidente == 2 {
				isMutante++
			}
		} else {
			isConcidente = 0
		}
	}
	<-c
	return isMutante
}

func Stats(w http.ResponseWriter, r *http.Request) {
	informe, err := getRatio()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Se presento un error")
		return
	}

	response, err := informe.ToJson()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
