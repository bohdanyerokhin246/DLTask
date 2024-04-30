package main

import "math"

type Block struct {
	Value uint64
}

func (b *Block) setHex(hex string) error {
	validatedHex, err := validateHex(hex)
	if err != nil {
		return err
	}

	value := *new(uint64)
	for i, r := range validatedHex {
		pow := math.Pow(16, float64(len(validatedHex)-i-1))
		if r >= '0' && r <= '9' {
			left := uint64(r - '0')
			value += left * uint64(pow)
		} else {
			left := uint64(uint64(r-'a') + 10)
			value += left * uint64(pow)
		}
	}
	b.setDecimal(value)

	return nil
}

func (b *Block) getHex() (hex string) {
	value := b.getDecimal()
	for value > 0 {
		hex = string(hexDigits[value%16]) + hex
		value /= 16
	}
	return
}

func (b *Block) getDecimal() uint64 {
	return b.Value
}

func (b *Block) setDecimal(value uint64) {
	b.Value = value
}

func (b *Block) ADD(other Block) Block {
	return Block{b.getDecimal() + other.getDecimal()}
}
