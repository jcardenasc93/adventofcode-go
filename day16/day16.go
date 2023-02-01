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
	packets := list.New()
	packets.PushBack(binary)
	findVersions(packets)
	fmt.Println(versionSum)
}

var versionSum int

func findVersions(packets *list.List) {
	literal := 4
	var left string

	for packets.Len() > 0 {
		front := packets.Front()
		bin := front.Value.(string)
		packets.Remove(front)
		if !(onlyZeros(bin)) {
			if len(bin) > 6 {
				_, tID := getPacketVType(bin)
				if tID == literal {
					left = handleLiteral(bin)
				} else {
					left = handleOperator(bin)
				}
			}
			if len(left) > 6 {
				packets.PushBack(left)
				findVersions(packets)
			}
		}
	}
}

func getPacketVType(bin string) (pv, tID int) {
	pv = parse3Bits(bin, 0, 3)
	versionSum += pv
	tID = parse3Bits(bin, 3, 6)
	return pv, tID
}

func parse3Bits(s string, start int, end int) int {
	packetV := s[start:end]
	v, _ := strconv.ParseInt(packetV, 2, 8)
	return int(v)
}

func handleLiteral(literal string) (left string) {
	i := 6
	for string(literal[i]) != "0" {
		i += 5
	}
	i += 5
	left = literal[i:]

	return left
}

func handleOperator(operator string) (left string) {
	if string(operator[6]) == "0" {
		left = operator[7+15:]
	} else {
		left = operator[7+11:]
	}

	return left
}

func getSubPacksBitsLen(bin string) int {
	bitsLen, _ := strconv.ParseUint(bin[7:7+15], 2, 8)
	return int(bitsLen)
}

func getSubPacksNum(bin string) int {
	packs, _ := strconv.ParseUint(bin[7:7+11], 2, 8)
	return int(packs)
}

func onlyZeros(bin string) bool {
	for _, c := range bin {
		if string(c) != "0" {
			return false
		}
	}
	return true
}
