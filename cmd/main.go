package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

const HEX = "ffffffffffff0000000000000800450000480001000040297c8a7f0000017f00000160000000000c2f400000000000000000000000000000000100000000000000000000000000000001000000000035003500080000"

// ConvertText2PcapFile take filename for input and output and base on this information
// try to convert text file into binary, using 'text2pcap'.
func ConvertText2PcapFile(iFilename string, oFilename string) {
	text2pcapCommand := fmt.Sprintf("text2pcap %s %s", iFilename, oFilename)
	_, err := exec.Command("/bin/sh", "-c", text2pcapCommand).Output()
	if err != nil {
		log.Fatal(err)
	}
}

// Hex2Array accept a hex in string and transform it into string list
// IN: var hex string = "00aabb"
// return ["00", "aa", "bb"]
func Hex2Array(hex string) (combinedHexArray []string) {
	var hexArray []string = strings.Split(hex, "")

	for i := 0; i < len(hexArray); i += 2 {
		combinedHexArray = append(combinedHexArray, hexArray[i]+hexArray[i+1])
	}
	return
}

// DumpHex transform a string HEX value into text
// format acceptable by the text2pcap
func DumpHex(hex string) (hexdump []string) {
	var hexArray []string = Hex2Array(hex)

	// TODO this code can be much simplier
	var line int
	var tmp_array []string

	var arrayLen int = len(hexArray)

	for i, s := range hexArray {
		if i%16 == 0 {
			if line != 0 {
				hexdump = append(hexdump, strings.Join(tmp_array, " "))
			}
			tmp_array = []string{fmt.Sprintf("%03d0", line)}
			line++
		}
		tmp_array = append(tmp_array, s)

		if i == arrayLen-1 {
			hexdump = append(hexdump, strings.Join(tmp_array, " "))
		}
	}
	return
}

func WriteDumpHex(hex string, filename string) {
	joinedDecodedHex := strings.Join(DumpHex(hex), "\n")

	err := os.WriteFile(filename, []byte(joinedDecodedHex), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

type TSharkOutputSource struct {
	Source map[string]interface{} `json:"_source"`
}

func DecodeTShark(filename string) string {
	output, err := exec.Command("/bin/sh", "-c", fmt.Sprintf("tshark -x -r %s -T json", filename)).Output()
	if err != nil {
		log.Fatal(err)
	}
	tsharkOutput := []TSharkOutputSource{}
	if err := json.Unmarshal(output, &tsharkOutput); err != nil {
		log.Fatal(err)
	}
	return string(output)
}

func DecodePacket(hex string) string {
	timeNow := fmt.Sprint(time.Now().UTC().UnixNano())
	inputFilename, outputFilename := fmt.Sprintf("%s.txt", timeNow), fmt.Sprintf("%s.bin", timeNow)

	WriteDumpHex(hex, inputFilename)
	ConvertText2PcapFile(inputFilename, outputFilename)

	defer cleanUp([]string{inputFilename, outputFilename})
	return DecodeTShark(outputFilename)
}

func cleanUp(filenames []string) {
	for _, x := range filenames {
		if err := os.Remove(x); err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	output := DecodePacket(HEX)
	fmt.Print(output)
}
