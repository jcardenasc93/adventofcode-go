package utils

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

func ParseStringToIntSlice(input string) []int {
	var output []int
	for _, items := range strings.Split(input, ",") {
		item, err := strconv.Atoi(items)
		CheckError(err)
		output = append(output, item)
	}
	return output
}

func ParseStrInt(number string) int {
	num, err := strconv.Atoi(number)
	CheckError(err)
	return num
}

func IsItemIn[T comparable](slice []T, item T) bool {
	for _, i := range slice {
		if item == i {
			return true
		}
	}
	return false
}

func DelItem[T comparable](slice []T, item T) []T {
	resp := []T{}
	for j, i := range slice {
		if item == i {
			resp = append(resp, slice[:j]...)
			resp = append(resp, slice[j+1:]...)
		}
	}
	return resp
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsMapIn[T comparable](mapSlice *[]map[string]T, mA map[string]T) bool {
	for _, m := range *mapSlice {
		if reflect.DeepEqual(m, mA) {
			return true
		}
	}
	return false
}
