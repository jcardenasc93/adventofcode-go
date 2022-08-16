package day3

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunDay3() {
	fmt.Println("Input diagnostic report:")
	// decodeDiagnostic()
	decodeDiagnosticV2()
}

type diagnosticDecoder struct {
	bits    map[int][]string
	count   map[int]map[string]int
	bitsqty int
}

func decodeDiagnostic() {
	var input string
	var powerDecoder diagnosticDecoder
	var powerComp int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		input = scanner.Text()

		if input == "" {
			break
		}
		powerComp = calcPowerConsumption(input, &powerDecoder)
	}
	fmt.Println("Power consumption is:", powerComp)
}

func decodeDiagnosticV2() {
	var lifeRating int
	var input string
	var decoder diagnosticDecoder
	lastInput := false
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		input = scanner.Text()

		if input == "" {
			lastInput = true
			lifeRating = calcLifeSupRating(input, &decoder, lastInput)
			break
		}
		calcLifeSupRating(input, &decoder, lastInput)
	}
	fmt.Println(lifeRating)
	fmt.Println("end")
}

func calcLifeSupRating(number string, decoder *diagnosticDecoder, last bool) int {
	if decoder.bits == nil {
		decoder.bits = map[int][]string{}
	}
	if decoder.count == nil {
		decoder.count = map[int]map[string]int{}
	}
	if number != "" {
		for i, bit := range strings.Split(number, "") {
			decoder.bits[i] = append(decoder.bits[i], string(bit))
			decoder.count[i] = countBits(decoder.bits[i])
		}
	}
	if last {
		decoder.bitsqty = len(decoder.bits)
		// Start calc
		o2Gen := calcO2Gen(decoder)
		cO2Scrub := calcCO2Scrub(decoder)
		fmt.Println(o2Gen, cO2Scrub)
		o2, _ := strconv.ParseInt(o2Gen, 2, 0)
		cO2, _ := strconv.ParseInt(cO2Scrub, 2, 0)
		fmt.Println(o2, cO2)
		return int(o2 * cO2)
	}
	return 0

}

func calcO2Gen(decoder *diagnosticDecoder) string {
	// Find max in each bit
	var o2GenBuf bytes.Buffer
	var o2Gen string
	var left []int
	var temp []int
	found := false
	for {
		if len(left) == 1 {
			break
		}
		for i := 0; i < len(decoder.count); i++ {
			max := getMaxScope(i, decoder, left)
			if len(left) > 1 {
				// Search only in left indexes
				for _, j := range left {
					// Search numbers that bit[i] == max
					if decoder.bits[i][j] == max {
						temp = append(temp, j)
					}
				}
				left = temp
				temp = []int{}
				if len(left) == 1 {
					found = true
				}
			} else {
				for j := 0; j < len(decoder.bits[0]); j++ {
					// Search numbers that bit[i] == max
					if decoder.bits[i][j] == max {
						left = append(left, j)
					}
				}
			}
			if found {
				break
			}
		}
	}
	for i := 0; i < decoder.bitsqty; i++ {
		o2GenBuf.WriteString(decoder.bits[i][left[0]])
	}
	o2Gen = o2GenBuf.String()
	return o2Gen
}

func calcCO2Scrub(decoder *diagnosticDecoder) string {
	// Find less in each bit
	var cO2GenBuf bytes.Buffer
	var cO2Gen string
	var left []int
	var temp []int
	found := false
	for {
		if len(left) == 1 {
			break
		}
		for i := 0; i < len(decoder.count); i++ {
			min := getMinScope(i, decoder, left)
			if len(left) > 1 {
				// Search only in left indexes
				for _, j := range left {
					// Search numbers that bit[i] == min
					if decoder.bits[i][j] == min {
						temp = append(temp, j)
					}
				}
				left = temp
				temp = []int{}
				if len(left) == 1 {
					found = true
				}
			} else {
				for j := 0; j < len(decoder.bits[0]); j++ {
					// Search numbers that bit[i] == min
					if decoder.bits[i][j] == min {
						left = append(left, j)
					}
				}
			}
			if found {
				break
			}
		}
	}
	for i := 0; i < decoder.bitsqty; i++ {
		cO2GenBuf.WriteString(decoder.bits[i][left[0]])
	}
	cO2Gen = cO2GenBuf.String()
	return cO2Gen
}

func calcPowerConsumption(number string, decoder *diagnosticDecoder) int {
	if decoder.bits == nil {
		decoder.bits = map[int][]string{}
	}
	if decoder.count == nil {
		decoder.count = map[int]map[string]int{}
	}
	for i, bit := range strings.Split(number, "") {
		decoder.bits[i] = append(decoder.bits[i], string(bit))
		decoder.count[i] = countBits(decoder.bits[i])
	}
	gamma, epsilon := calcGammaEpsilon(decoder)
	gammaInt, _ := strconv.ParseInt(gamma, 2, 0)
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 0)
	return int(gammaInt) * int(epsilonInt)

}

func countBits(bitsl []string) map[string]int {
	count := map[string]int{"0": 0, "1": 0}
	for _, bit := range bitsl {
		if bit == "0" {
			count["0"] = count["0"] + 1
		} else {
			count["1"] = count["1"] + 1
		}
	}
	return count
}

func calcGammaEpsilon(decoder *diagnosticDecoder) (string, string) {
	var gammastr bytes.Buffer

	// Find max in each bit
	for i := 0; i < len(decoder.count); i++ {
		max := getMax(decoder.count[i])
		gammastr.WriteString(max)
	}

	gamma := gammastr.String()
	epsilon := notBinString(gammastr.String())
	return gamma, epsilon
}

func getMax(count map[string]int) string {
	maxCount := 0
	var max string
	for k, v := range count {
		if v > maxCount {
			max = k
			maxCount = v
		}
	}
	return max
}

func getMaxScope(index int, decoder *diagnosticDecoder, remainingIdx []int) string {
	var max string
	maxCount := 0
	count := countBitsV2(index, remainingIdx, decoder)
	for k, v := range count {
		if v > maxCount {
			max = k
			maxCount = v
		}
		// In case of tie must return "1"
		if v == maxCount && max != k {
			max = "1"
		}
	}
	return max

}

func getMinScope(index int, decoder *diagnosticDecoder, remainingIdx []int) string {
	var min string
	minCount := 1000000
	count := countBitsV2(index, remainingIdx, decoder)

	for k, v := range count {
		if v < minCount {
			min = k
			minCount = v
		}
		// In case of tie must return "0"
		if v == minCount && min != k {
			min = "0"
		}
	}
	return min
}

func countBitsV2(index int, remainingIdx []int, decoder *diagnosticDecoder) map[string]int {
	count := map[string]int{"0": 0, "1": 0}
	if len(remainingIdx) > 0 {
		for _, idx := range remainingIdx {
			bit := decoder.bits[index][idx]
			if bit == "0" {
				count["0"] = count["0"] + 1
			} else {
				count["1"] = count["1"] + 1
			}
		}
	} else {
		for _, bit := range decoder.bits[index] {
			if bit == "0" {
				count["0"] = count["0"] + 1
			} else {
				count["1"] = count["1"] + 1
			}
		}
	}
	return count
}

func notBinString(num string) string {
	var result bytes.Buffer
	notOp := map[string]string{
		"0": "1",
		"1": "0",
	}
	for _, bit := range num {
		result.WriteString(notOp[string(bit)])
	}
	return result.String()

}
