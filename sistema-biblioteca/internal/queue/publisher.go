package queue

import (
	"encoding/json"
	"os"
	"github.com/rabbitmq/amqp091-go"
)

type EventoBiblioteca struct {
	Evento    string `json:"evento"`
	UsuarioID uint   `json:"usuario_id"`
	LivroID   uint   `json:"livro_id"`
	Mensagem  string `json:"mensagem"`
}

// PublishEvento envia um evento JSON para a fila biblioteca_queue
func PublishEvento(evento EventoBiblioteca) error {
	amqpURL := os.Getenv("RABBITMQ_URL")
	if amqpURL == "" {
		amqpURL = "amqp://guest:guest@localhost:5672/"
	}

	conn, err := amqp091.Dial(amqpURL)
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"biblioteca_queue",
		true, false, false, false, nil,
	)
	if err != nil {
		return err
	}

	body, err := json.Marshal(evento)
	if err != nil {
		return err
	}

	return ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
