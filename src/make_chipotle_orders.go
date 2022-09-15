package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/segmentio/kafka-go"
)

var MenuLen = len(chipotleMenu)

type CustomerInfo struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	City  string `json:"city"`
	State string `json:"state"`
}

type ChipotleOrder struct {
	Time     string       `json:"time"`
	Customer CustomerInfo `json:"customer"`
	Order    []Item       `json:"order"`
	Card     string       `json:"card"`
}

func createAge() int {
	return rand.Intn(76-14) + 14
}

func createCustomer() CustomerInfo {
	return CustomerInfo{gofakeit.Name(), createAge(), gofakeit.City(), gofakeit.State()}
}

func createOrder() []Item {
	itemCount := rand.Intn(2) + 1
	items := make([]Item, itemCount)
	for i := 0; i < itemCount; i++ {
		items[i] = chipotleMenu[rand.Intn(MenuLen)]
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

	orderStream := make(chan ChipotleOrder)

	go func() {
		defer close(orderStream)

		for {
			time.Sleep(30 * time.Millisecond)
			orderStream <- getOrder()
		}
	}()

	for order := range orderStream {
		fmt.Println(" order ==> ", order)
		orderJSON, err := json.Marshal(order)
		if err != nil {
			log.Println("failed to marshall order: ", err)
		}

		err = push(w, context.Background(), []byte(order.Time), orderJSON)
		if err != nil {
			log.Println("failed to write message: ", err)
		}
	}
}
