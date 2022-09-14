package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/segmentio/kafka-go"
)

type Customer struct {
	name string
	age  int
}

type Order struct {
	time     string
	customer Customer
	order    Item
	quantity int
}

func (o Order) String() string {
	age := strconv.Itoa(o.customer.age)
	price := fmt.Sprintf("%.2f", o.order.price)
	quantity := strconv.Itoa(o.quantity)

	return o.time + "," + o.customer.name + "," + age + "," + o.order.name + "," + price + "," + quantity
}

func createName() string {
	var n strings.Builder

	for i := 1; i < 6; i++ {
		r := rune(rand.Intn(122-97) + 97)
		if i == 1 {
			r = unicode.ToUpper(r)
		}
		n.WriteRune(r)
	}

	return n.String()
}

func createAge() int {
	return rand.Intn(76-14) + 14
}

func createCustomer() Customer {
	return Customer{createName(), createAge()}
}

func getQuantity() int {
	return rand.Intn(4-1) + 1
}

func getOrder() Order {
	return Order{
		time.Now().UTC().Format("2006-01-02 15:04:05"),
		createCustomer(),
		chipotleMenu[rand.Intn(len(chipotleMenu))],
		getQuantity(),
	}
}

func main() {
	for {
		w := &kafka.Writer{
			Addr:  kafka.TCP("localhost:9092"),
			Topic: "orders",
		}
		ctx := context.Background()
		got := getOrder()
		err := w.WriteMessages(
			ctx,
			kafka.Message{
				Key:   []byte(got.time),
				Value: []byte(got.String()),
			},
		)

		if err != nil {
			log.Fatal("failed to write messages: ", err)
		}

		if err := w.Close(); err != nil {
			log.Fatal("failed to close writer: ", err)
		}
	}
}
