package main

import (
	"fmt"
	"log"

	"github.com/PacketHelper/goshark/v2/goshark"
)

func main() {
	fmt.Print("Starting goshark api...")

	router := goshark.CreateRouter()
	router.Run()
	if err := router.Run(); err != nil {
		log.Fatalf("cannot start http server %s", err)
	}
}
