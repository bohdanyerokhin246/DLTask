package main

import (
	"fmt"
	"math"
	"strings"
)

const hexDigits string = "0123456789abcdef"

func reverseString(s string) string {
	runeArray := []rune(s)
	for i := 0; i < len(runeArray)/2; i++ {
		runeArray[i], runeArray[len(runeArray)-i-1] = runeArray[len(runeArray)-i-1], runeArray[i]
	}
	return string(runeArray)
}

func breakStringIntoBlocks(input string, blockSize int) []string {
	numberOfBlocks := uint(math.Ceil(float64(len(input)) / float64(blockSize)))
	blocks := make([]string, numberOfBlocks)
	reversedInput := reverseString(input)

	for i := 0; i < int(numberOfBlocks); i++ {
		if i*blockSize+blockSize > len(reversedInput) {
			blocks[i] = reverseString(reversedInput[i*blockSize:])
		} else {
			blocks[i] = reverseString(reversedInput[i*blockSize : i*blockSize+blockSize])
		}
	}

	return blocks
}

func validateHex(hex string) (string, error) {
	hex = strings.ToLower(hex)
	if len(hex) > 16 {
		return "", fmt.Errorf("max hex length is 16")
	}
	for _, r := range hex {
		if !strings.ContainsRune(hexDigits, r) {
			return "", fmt.Errorf("'%s' is not a hex digit", string(r))
		}
	}
	return hex, nil
}

func AddLeadingZeros(value string, size int) string {
	missingZerosCount := size - len(value)
	var sb strings.Builder
	for i := 0; i < missingZerosCount; i++ {
		sb.WriteString("0")
	}
	sb.WriteString(value)
	return sb.String()
}
