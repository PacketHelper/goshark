package main

import (
	"fmt"

	"github.com/PacketHelper/goshark/v2/goshark"
)

func main() {
	fmt.Print("Starting goshark api...")
	goshark.HttpServer()
}
