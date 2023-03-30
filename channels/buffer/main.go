package main

import "fmt"

func main() {

	//Declarando canal
	ch := make(chan string, 4)

	//Enviando mensagens para o canal

	// canal[3][2][1] --> output do canal

	ch <- "msg1"
	fmt.Println(<-ch)
	ch <- "msg2"
	ch <- "msg3"
	ch <- "msg4"
	ch <- "msg5"

	//Channel implementa lógica FIFO = First in First out.
	fmt.Println(ch)

	close(ch)

}
