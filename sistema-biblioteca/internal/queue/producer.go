
package main

import (
	"log"
	"os"
	"github.com/rabbitmq/amqp091-go"
)

// PublishMessage envia uma mensagem para a fila RabbitMQ

func main() {
   amqpURL := os.Getenv("RABBITMQ_URL")
   if amqpURL == "" {
	   amqpURL = "amqp://guest:guest@localhost:5672/"
   }

   conn, err := amqp091.Dial(amqpURL)
   if err != nil {
	   log.Fatalf("erro ao conectar no RabbitMQ: %v", err)
   }
   defer conn.Close()

   ch, err := conn.Channel()
   if err != nil {
	   log.Fatalf("erro ao abrir canal: %v", err)
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
	   log.Fatalf("erro ao declarar fila: %v", err)
   }

   msg := "Olá, biblioteca! Mensagem enviada para a fila."
   err = ch.Publish(
	   "",     // exchange
	   q.Name, // routing key (nome da fila)
	   false,  // mandatory
	   false,  // immediate
	   amqp091.Publishing{
		   ContentType: "text/plain",
		   Body:        []byte(msg),
	   },
   )
   if err != nil {
	   log.Fatalf("erro ao publicar mensagem: %v", err)
   }

   log.Printf("Mensagem publicada: %s", msg)
}
