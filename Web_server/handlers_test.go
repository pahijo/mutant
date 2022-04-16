package main

import (
	//"fmt"
	"sync"
	"testing"
)

func TestIsValid(t *testing.T) {
	var mutant DnaList
	mutant.AgregarLista("ATGCGA")
	mutant.AgregarLista("CAGTGC")
	mutant.AgregarLista("TTATGT")
	mutant.AgregarLista("AGAAGG")
	mutant.AgregarLista("CCCCTA")
	mutant.AgregarLista("TCACTG")
	//fmt.Println(mutant)
	_, isValido, message := isValid(mutant, 6)
	//fmt.Println(matrizDna)
	if !isValido {
		t.Error(message)
	}
}

func TestIsMutantV(t *testing.T) {

	var mutant DnaList
	mutant.AgregarLista("ATGCGA")
	mutant.AgregarLista("CAGTGC")
	mutant.AgregarLista("TTATGT")
	mutant.AgregarLista("AGAAGG")
	mutant.AgregarLista("CCCCTA")
	mutant.AgregarLista("TCACTX")
	matrizDna, _, _ := isValid(mutant, 6)
	isMutante := 0
	c := make(chan int, 3)
	var wg sync.WaitGroup
	c <- 1
	wg.Add(1)
	//go func() {
	isMutante = isMutante + isMutantV(matrizDna, &wg, c)
	//}
	wg.Wait()
	if isMutante != 1 {
		t.Error("Prueba no superada")
	}
}

func TestIsMutantH(t *testing.T) {

	var mutant DnaList
	mutant.AgregarLista("ATGCGA")
	mutant.AgregarLista("CAGTGC")
	mutant.AgregarLista("TTATGT")
	mutant.AgregarLista("AGAAGG")
	mutant.AgregarLista("CCCCTA")
	mutant.AgregarLista("TCACTX")
	matrizDna, _, _ := isValid(mutant, 6)
	isMutante := 0
	c := make(chan int, 3)
	var wg sync.WaitGroup
	c <- 1
	wg.Add(1)
	//go func() {
	isMutante = isMutante + isMutantH(matrizDna, &wg, c)
	//}
	wg.Wait()
	if isMutante != 1 {
		t.Error("Prueba no superada")
	}
}

func TestIsMutantO(t *testing.T) {

	var mutant DnaList
	mutant.AgregarLista("ATGCGA")
	mutant.AgregarLista("CAGTGC")
	mutant.AgregarLista("TTATGT")
	mutant.AgregarLista("AGAAGG")
	mutant.AgregarLista("CCCCTA")
	mutant.AgregarLista("TCACTX")
	matrizDna, _, _ := isValid(mutant, 6)
	isMutante := 0
	c := make(chan int, 3)
	var wg sync.WaitGroup
	c <- 1
	wg.Add(1)
	//go func() {
	isMutante = isMutante + isMutantO(matrizDna, &wg, c)
	//}
	wg.Wait()
	if isMutante != 1 {
		t.Error("Prueba no superada")
	}
}
