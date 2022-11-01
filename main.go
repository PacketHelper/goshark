package main

import (
	"fmt"

	"github.com/PacketHelper/goshark/v2/goshark"
)

func main() {
	test()
	// fmt.Print("Starting goshark api...")
	goshark.HttpServer()
}

func test() {
	s := "00001Cffffff0000000000000800450000340001000040047cc37f0000017f0000014500002000010000402f7cac7f0000017f000001000000000035003500080000"

	p := goshark.DecodePacketXML(s, false)
	fmt.Print(p)
}
