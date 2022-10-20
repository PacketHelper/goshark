package goshark

import (
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
)

const TSharkHeader = "\xd4\xc3\xb2\xa1\x02\x00\x04\x00\x00\x00\x00\x00\x00\x00\x00\x00\xff\x7f\x00\x00\x01\x00\x00\x00"
const TWO = "\x91\xbeFc\ng\r\x00^\x00\x00\x00^\x00\x00\x00"
const THR = "\xff\xff\xff\xff\xff\xff\x00\x00\x00\x00\x00\x00\x08\x00E\x00\x00P\x00\x01\x00\x00@)|\x82\x7f\x00\x00\x01\x7f\x00\x00\x01`\x00\x00\x00\x00\x14\x06@\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01\x00\x14\x00P\x00\x00\x00\x00\x00\x00\x00\x00P\x02 \x00\x8f}\x00\x00"

func CreateTSharkHeader() (header string) {
	// var _bytes []byte

	// _bytes = append(_bytes, []byte("1")...)
	// fmt.Println(_bytes)

	// a := fmt.Sprintf("%b", 1)
	// a, _ := hex.DecodeString("7fff")
	// b, _ := strconv.ParseInt("3", 2, 0)
	// fmt.Print(b)
	// fmt.Println([]byte(a))
	fmt.Println([]byte(TSharkHeader))
	fmt.Println(hex.DecodeString("ff"))
	return header
}

func createTSharkPacketInformationHeader() {

}

// RunTSharkProcess spawn a TShark process and transfer to it packets to decode
func RunTSharkProcess() {
	subProcess := exec.Command(
		"tshark", "-l", "-n", "-T", "pdml", "-i", "-",
	)

	stdin, err := subProcess.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		// stdin.Write([]byte{1, 67})
		// v, _ := hex.DecodeString(ONE)
		// v2, _ := hex.DecodeString(TWO)
		// v3, _ := hex.DecodeString(THR)

		// io.WriteString(stdin, ONE)
		io.WriteString(stdin, TWO)
		io.WriteString(stdin, THR)
		// stdin.Write(v3)
		// io.WriteString(stdin, fmt.Sprintf("%s%s%s", ONE, TWO, THR))
	}()

	out, err := subProcess.CombinedOutput()
	if err != nil {
		log.Fatal("something went wrong with decoding packet...")
	}

	out2 := string(out)
	start, stop := strings.Index(out2, "<packet>"), strings.Index(out2, "</packet>")
	fmt.Printf("\n%s\n", out[start:stop+9])

	var packet Packet
	xml.Unmarshal(out[start:stop+9], &packet)
	fmt.Println(packet)
}
