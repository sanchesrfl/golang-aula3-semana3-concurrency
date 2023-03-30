package main

import (
	"fmt"
	"time"
)

//estrutura que modela a mensagem do serviço de mensageria
type Message struct {
	ID        int
	Content   string
	CreatedAt time.Time //timestamp
}

func main() {

	//primeiro canal que recebe mensagens
	receiver := make(chan Message, 20)

	//segundo canal que é repositório a ser lido
	sender := make(chan Message, 20)

	//rotina que transmite esses dados entre canais
	go func() {
		for msg := range receiver {
			time.Sleep(2 * time.Second)
			sender <- msg
		}
	}()

	//rotina que envia as mensagens
	go func() {
		for msg := range sender {

			fmt.Println("A mensagem recebida é: ", msg.Content)
			fmt.Println("Seu timestamp é: ", msg.CreatedAt)
			time.Sleep(2 * time.Second)
		}
	}()

	//rotina que envia mensagens da aplicação para o canal de recebimento
	for i := 1; i <= 20; i++ {
		msg := Message{
			ID:        i,
			Content:   fmt.Sprintf("Mensagem %d recebida.", i),
			CreatedAt: time.Now(),
		}
		//mensagem sendo enviada ao canal recebedor
		receiver <- msg
		time.Sleep(2 * time.Second)
	}

	time.Sleep(4 * time.Second)
	close(receiver)
	close(sender)

}
