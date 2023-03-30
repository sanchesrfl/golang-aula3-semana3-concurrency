package main

/*
Go Routines
Implementação Geral/Básica
*/

import (
	"fmt"
	"time"
)

func Hello() {
	fmt.Println("Hello")
	time.Sleep(time.Second * 3)
}

//também é uma Go routine.
func main() {

	//comando Go inicia uma nova rotina.
	go Hello()

	//segunda Go rotina.
	go func() {
		for i := 1; i <= 4; i++ {
			fmt.Println("Fluxo de dados batch number: ", i)
			time.Sleep(time.Second * 2)
		}
	}()

	time.Sleep(time.Second * 10)
	fmt.Println("Processamento executado com sucesso.")

}
