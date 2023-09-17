package main

import (
	"context"
	"fmt"

	"github.com/an33k/go-microservice/app"
)

func main() {
	app := app.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}