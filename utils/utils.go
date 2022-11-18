package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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
