package main

import (
	"fmt"

	"github.com/PacketHelper/goshark/v2/gosharkapi"
)

func main() {
	fmt.Print("Starting goshark api...")
	gosharkapi.HttpServer()
}
