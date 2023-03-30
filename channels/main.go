package main

/*
Implementar Canais (Channels)
*/

import "fmt"

func main() {

	//criar um canal sem buffer
	ch := make(chan string)

	//enviando mensagens para o canal
	go func() {
		ch <- "Mensagem Inicial"
	}()

	//recuperando msg do canal
	msg := <-ch
	fmt.Println(msg)

	//fechando canal
	close(ch)

}
