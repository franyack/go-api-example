package main

import (
	"go-api-example/internal/infraestructure/delivery/webapi"
)

func main() {
	deliveryStrategy := webapi.New()
	deliveryStrategy.Start()
}
