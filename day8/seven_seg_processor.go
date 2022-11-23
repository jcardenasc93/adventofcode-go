package day8

import (
	"strings"
)

// This var maps seven segment display segments as follows:
//  000
// 5   1
// 5   1
//  666
// 4   2
// 4   2
//  333
// So each number has a value based on the sum of the segments needed
// to display the number.
// Ex: To dispaly '1' segments 1 and 2 should be on so 1+2 = 3

var sevenSeg = map[int][]int{
	0: {15, 6},
	1: {3, 2},
	2: {14, 5},
	3: {12, 5},
	4: {14, 4},
	5: {16, 5},
	6: {20, 6},
	7: {3, 3},
	8: {21, 7},
	9: {17, 6},
}

type SevenSegProcessor struct {
	segmentLetters []string
	decipher       map[int]string
}

func sevenSegDecipher(input []string) *SevenSegProcessor {
	processor := SevenSegProcessor{
		segmentLetters: []string{"a", "b", "c", "d", "e", "f", "g"},
		decipher:       make(map[int]string),
	}
	// First look for unique cases; '1', '7', '4'
	var signal1 string
	var signal7 string
	var signal4 string
	var criticalSignals []string
	for _, s := range input {
		if len(s) == 2 {
			signal1 = s
		}
		if len(s) == 3 {
			signal7 = s
		}
		if len(s) == 4 {
			signal4 = s
		}
		if len(s) == 5 {
			criticalSignals = append(criticalSignals, s)
		}
	}
	// map segments for '1'
	processor.updateDecipher(1, string(signal1[0]))
	processor.updateDecipher(2, string(signal1[1]))

	// map segment for '7'
	for i := 0; i < len(signal7); i++ {
		s := string(signal7[i])
		if strings.Contains(signal1, s) == false {
			processor.updateDecipher(0, s)
		}
	}
	processor.mapSegment6(&signal4, &criticalSignals)
	processor.mapSegment5(&signal4)
	processor.mapSegment3(&criticalSignals)
	// map missing letter
	processor.updateDecipher(4, processor.segmentLetters[0])
	return &processor
}

func (processor *SevenSegProcessor) updateDecipher(key int, value string) {
	processor.decipher[key] = value
	processor.removeLetter(value)
}

func (processor *SevenSegProcessor) removeLetter(letter string) {
	for i, l := range processor.segmentLetters {
		if l == letter {
			processor.segmentLetters = append(processor.segmentLetters[:i], processor.segmentLetters[i+1:]...)
			break
		}
	}
}

func (processor *SevenSegProcessor) mapSegment6(signal4 *string, otherSignals *[]string) {
	// Look for a common char between signal for '4' and all signals with len = 5
	var common string
	for i := 0; i < len(*signal4); i++ {
		s := string((*signal4)[i])
		for _, signal := range *otherSignals {
			if strings.Contains(signal, s) == true {
				if common == "" {
					common = s
				} else if common != s {
					common = ""
					break
				}
			} else {
				common = ""
				break
			}
		}
		if common != "" {
			break
		}
	}
	processor.updateDecipher(6, common)
}

func (processor *SevenSegProcessor) getDecipherKeys() []string {
	currentKeys := make([]string, len(processor.decipher))
	i := 0
	for _, v := range processor.decipher {
		currentKeys[i] = v
		i += 1
	}
	return currentKeys
}

func (processor *SevenSegProcessor) mapSegment5(signal4 *string) {
	// Segment 5 is the left letter not present in decipher map
	isPresent := true
	keys := processor.getDecipherKeys()
	for _, s := range *signal4 {
		ss := string(s)
		for _, k := range keys {
			if ss == k {
				isPresent = true
				break
			} else {
				isPresent = false
			}
		}
		if isPresent == false {
			processor.updateDecipher(5, ss)
			break
		}
	}
}

// Look for a common char between all signals with len == 5 but not same char assigned to segment 6 nor segment 0
func (processor *SevenSegProcessor) mapSegment3(otherSignals *[]string) {
	var common string
	signal := (*otherSignals)[0]
	for i := 0; i < len(signal); i++ {
		s := string((signal)[i])
		if s != processor.decipher[6] && s != processor.decipher[0] {
			for _, signal := range *otherSignals {
				if strings.Contains(signal, s) == true {
					if common == "" {
						common = s
					} else if common != s {
						common = ""
						break
					}
				} else {
					common = ""
					break
				}
			}
			if common != "" {
				break
			}
		}
	}
	processor.updateDecipher(3, common)
}

func (processor *SevenSegProcessor) decodeValue(signal string) int {
	var sum int
	if len(signal) == 7 {
		return 8
	}
	if len(signal) == 2 {
		return 1
	}
	if len(signal) == 3 {
		return 7
	}
	if len(signal) == 4 {
		return 4
	}
	for _, s := range signal {
		ss := string(s)
		for k, v := range processor.decipher {
			if ss == v {
				sum += int(k)
				break
			}
		}
	}
	if sum == 14 {
		if len(signal) == 4 {
			return 4
		}
		return 2
	}
	for digit, val := range sevenSeg {
		if val[0] == sum && len(signal) == val[1] {
			return digit
		}
	}
	// If reach this point means that segments 1, 2 are trucated
	processor.decipher[1], processor.decipher[2] = processor.decipher[2], processor.decipher[1]
	return processor.decodeValue(signal)
}

func (processor *SevenSegProcessor) getCipherNum(stringVals []string) int {
	var number int
	multiplier := 1000
	for _, val := range stringVals {
		digit := processor.decodeValue(val)
		number += multiplier * digit
		multiplier = multiplier / 10
	}
	return number
}
