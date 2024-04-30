package main

type BigInt struct {
	blocks []Block
}

func (bi *BigInt) setHex(hex string) error {
	return bi.setValue(16, hex, func(s string) (Block, error) {
		var b Block
		if err := b.setHex(s); err != nil {
			return Block{0}, err
		}
		return b, nil
	})
}

func (bi *BigInt) GetHex() (hex string) {
	return bi.getValue(16, func(b Block) string {
		return b.getHex()
	})
}

func (bi *BigInt) setValue(blockSize int, value string, setter func(string) (Block, error)) error {
	inputBlocks := breakStringIntoBlocks(value, blockSize)
	resultBlocks := make([]Block, 0)
	for _, block := range inputBlocks {
		u, err := setter(block)
		if err != nil {
			return err
		}
		resultBlocks = append(resultBlocks, u)
	}
	bi.setBlocks(resultBlocks)
	return nil
}

func (bi *BigInt) getValue(blockSize int, getter func(block Block) string) (result string) {
	for i, block := range bi.getBlocks() {
		blockValue := getter(block)
		if i != len(bi.getBlocks())-1 {
			blockValue = AddLeadingZeros(blockValue, blockSize)
		}
		result = blockValue + result
	}
	return
}

func (bi *BigInt) setBlocks(blocks []Block) {
	bi.blocks = blocks
}

func (bi *BigInt) getBlocks() []Block {
	return bi.blocks
}

func (bi *BigInt) appendBlock(block Block) {
	bi.setBlocks(append(bi.getBlocks(), block))
}

func (bi *BigInt) ADD(other BigInt) (result BigInt) {
	carry := Block{0}
	thisBlocks := bi.getBlocks()
	otherBlocks := other.getBlocks()
	for i := 0; i < len(thisBlocks) || i < len(otherBlocks); i++ {
		if i >= len(thisBlocks) {
			result.appendBlock(otherBlocks[i].ADD(carry))
			carry = Block{0}
		} else if i >= len(otherBlocks) {
			result.appendBlock(thisBlocks[i].ADD(carry))
			carry = Block{0}
		} else {
			sum := thisBlocks[i].ADD(otherBlocks[i])
			sum = sum.ADD(carry)
			result.appendBlock(sum)
			if sum.getDecimal() < thisBlocks[i].getDecimal() || sum.getDecimal() < otherBlocks[i].getDecimal() {
				carry = Block{1}
			} else {
				carry = Block{0}
			}
		}
	}
	if carry.getDecimal() > 0 {
		result.appendBlock(carry)
	}
	return
}

func (bi *BigInt) moreThan(other BigInt) bool {
	if len(bi.blocks) > len(other.blocks) {
		return true
	} else if len(bi.blocks) == len(other.blocks) {
		comp := 0
		for i, _ := range bi.blocks {
			if bi.blocks[i].Value > other.blocks[i].Value {
				comp += 1
			} else if bi.blocks[i] == other.blocks[i] {
				comp += 0
			} else {
				comp -= 1
			}
		}
		if comp > 0 {
			return true
		} else {
			return false
		}
	}
	return false
}

func (bi *BigInt) moreOrEqualThan(other BigInt) bool {
	if len(bi.blocks) > len(other.blocks) {
		return true
	} else if len(bi.blocks) == len(other.blocks) {
		comp := 0
		for i, _ := range bi.blocks {
			if bi.blocks[i].Value > other.blocks[i].Value {
				comp += 1
			} else if bi.blocks[i].Value == other.blocks[i].Value {
				comp += 0
			} else {
				comp -= 1
			}
		}
		if comp >= 0 {
			return true
		} else {
			return false
		}
	}
	return false
}

func (bi *BigInt) lessThan(other BigInt) bool {
	if len(bi.blocks) < len(other.blocks) {
		return true
	} else if len(bi.blocks) == len(other.blocks) {
		comp := 0
		for i, _ := range bi.blocks {
			if bi.blocks[i].Value < other.blocks[i].Value {
				comp += 1
			} else if bi.blocks[i].Value == other.blocks[i].Value {
				comp += 0
			} else {
				comp -= 1
			}
		}
		if comp > 0 {
			return true
		} else {
			return false
		}
	}
	return false
}

func (bi *BigInt) lessOrEqualThan(other BigInt) bool {
	if len(bi.blocks) < len(other.blocks) {
		return true
	} else if len(bi.blocks) == len(other.blocks) {
		comp := 0
		for i, _ := range bi.blocks {
			if bi.blocks[i].Value < other.blocks[i].Value {
				comp += 1
			} else if bi.blocks[i].Value == other.blocks[i].Value {
				comp += 0
			} else {
				comp -= 1
			}
		}
		if comp >= 0 {
			return true
		} else {
			return false
		}
	}
	return false
}

func (bi *BigInt) equal(other BigInt) bool {
	if len(bi.blocks) < len(other.blocks) {
		return false
	} else if len(bi.blocks) == len(other.blocks) {
		comp := 0
		for i, _ := range bi.blocks {
			if bi.blocks[i].Value == other.blocks[i].Value {
				comp += 1
			} else {
				return false
			}
		}
		if comp > 0 {
			return true
		}
	}
	return false
}
