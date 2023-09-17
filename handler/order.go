package handler

import (
	"fmt"
	"net/http"
)

type Order struct {

}

func(o *Order) CreateOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create an order")
}

func(o *Order) ListOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("list all order")
}

func(o *Order) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get an ordern by ID")
}

func(o *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updates an ordern by ID")
}

func(o *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deletes an ordern by ID")
}