package main

import (
	"fmt"
	"log"

	"github.com/PacketHelper/goshark/v2/goshark"
	ps "github.com/mitchellh/go-ps"
)

func findPyShark() {
	processes, err := ps.Processes()
	if err != nil {
		log.Fatal(err)
	}

	var processObj ps.Process
	for x := range processes {
		p := processes[x]
		processName := p.Executable()
		if processName == "tshark" {
			processObj = p
			break
		}
	}
	fmt.Print(processObj.PPid())
}

func main() {
	fmt.Print("Starting goshark api...")
	// goshark.HttpServer()
	// goshark.RunTSharkProcess()
	goshark.CreateTSharkHeader()
}
