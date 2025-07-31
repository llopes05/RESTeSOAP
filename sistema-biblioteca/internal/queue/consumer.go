package main

import (
	"fmt"
	"log"
	"os"
	"github.com/rabbitmq/amqp091-go"
)

func main() {
	amqpURL := os.Getenv("RABBITMQ_URL")
	if amqpURL == "" {
		amqpURL = "amqp://guest:guest@localhost:5672/"
	}

   conn, err := amqp091.Dial(amqpURL)
	if err != nil {
		log.Fatalf("Erro ao conectar no RabbitMQ: %v", err)
	}
	defer conn.Close()

   ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Erro ao abrir canal: %v", err)
	}
	defer ch.Close()

   q, err := ch.QueueDeclare(
	   "biblioteca_queue", // nome da fila
	   true,   // durável
	   false,  // auto-delete
	   false,  // exclusiva
	   false,  // no-wait
	   nil,    // argumentos
   )
	if err != nil {
		log.Fatalf("Erro ao declarar fila: %v", err)
	}

   msgs, err := ch.Consume(
	   q.Name, // queue
	   "",    // consumer
	   true,   // auto-ack
	   false,  // exclusive
	   false,  // no-local
	   false,  // no-wait
	   nil,    // args
   )
	if err != nil {
		log.Fatalf("Erro ao registrar consumidor: %v", err)
	}

	log.Println("Aguardando mensagens. Para sair pressione CTRL+C")
   for msg := range msgs {
	   fmt.Printf("Mensagem recebida: %s\n", msg.Body)
   }
}
