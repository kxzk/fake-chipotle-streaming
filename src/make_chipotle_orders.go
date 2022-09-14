package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/segmentio/kafka-go"
)

type CustomerInfo struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	City  string `json:"city"`
	State string `json:"state"`
}

type ChipotleOrder struct {
	Time       string       `json:"time"`
	Customer   CustomerInfo `json:"customer"`
	Order      []Item       `json:"order"`
	CreditCard string       `json:"creditCard"`
}

func createAge() int {
	return rand.Intn(76-14) + 14
}

func createCustomer() CustomerInfo {
	return CustomerInfo{gofakeit.Name(), createAge(), gofakeit.City(), gofakeit.State()}
}

func createOrder() []Item {
	itemCount := rand.Intn(3) + 1
	items := make([]Item, itemCount)
	for i := 0; i < itemCount; i++ {
		items = append(items, chipotleMenu[rand.Intn(len(chipotleMenu))])
	}
	return items
}

func getOrder() ChipotleOrder {
	return ChipotleOrder{
		time.Now().UTC().Format("2006-01-02 15:04:05"),
		createCustomer(),
		createOrder(),
		gofakeit.CreditCardType(),
	}
}

func push(writer *kafka.Writer, parent context.Context, key, value []byte) (err error) {
	message := kafka.Message{Key: key, Value: value}
	return writer.WriteMessages(parent, message)
}

func main() {
	w := &kafka.Writer{
		Addr:         kafka.TCP("localhost:9092"),
		Topic:        "orders",
		Async:        true,
		RequiredAcks: kafka.RequireNone,
	}
	defer w.Close()

	for {
		go func() {
			chipotleOrder := getOrder()
			orderMarshalled, err := json.Marshal(chipotleOrder)
			if err != nil {
				log.Println("failed to marshall chipotle order: ", err)
			}

			err = push(w, context.Background(), []byte(chipotleOrder.Time), orderMarshalled)
			if err != nil {
				log.Println("failed to write messages: ", err)
			}
		}()
	}
}
