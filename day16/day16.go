package day16

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"

	"github.com/jcardenasc93/adventofcode-go/utils"
)

func RunDay16() {
	readInputFile()
}

func readInputFile() {
	file, err := os.Open("file.txt")
	utils.CheckError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			processTransmission(line)
		}
	}
}

func processTransmission(l string) {
	hexaBin := map[string]string{
		"0": "0000",
		"1": "0001",
		"2": "0010",
		"3": "0011",
		"4": "0100",
		"5": "0101",
		"6": "0110",
		"7": "0111",
		"8": "1000",
		"9": "1001",
		"A": "1010",
		"B": "1011",
		"C": "1100",
		"D": "1101",
		"E": "1110",
		"F": "1111",
	}
	var binary string
	for _, hex := range l {
		binary = fmt.Sprintf("%s%s", binary, hexaBin[string(hex)])
	}
	fmt.Println(binary)
	packets := list.New()
	packets.PushBack(binary)
	findVersions(packets)
}

var versionSum int

func findVersions(packets *list.List) {
	literal := 4

	for packets.Len() > 0 {
		front := packets.Front()
		bin := front.Value.(string)
		packets.Remove(front)
		pv := parse3Bits(bin, 0, 3)
		versionSum += pv
		tID := parse3Bits(bin, 3, 6)
		if tID == literal {
			i := 6
			for string(bin[i]) != "0" {
				i += 5
			}
			i += 4
			left := bin[i:]
			if len(left) > 6 {
				packets.PushBack(bin[i:])
				findVersions(packets)
			}
		}
	}
	fmt.Println(versionSum)
}

func parse3Bits(s string, start int, end int) int {
	packetV := s[start:end]
	v, _ := strconv.ParseInt(packetV, 2, 8)
	return int(v)
}

// func getSubPacks(pack string) []string {
// 	subPacks := []string{}
// 	if string(pack[7]) == "0" {
// 		subPacksBits, _ := strconv.ParseUint(pack[8:24], 2, 8)
//
// 	}
// 	return subPacks
// }
