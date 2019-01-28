/*
 * Package bitarray implements the Sieve interface.
 * This is *NOT* a threadsafe package.
 */
package sieve

import (
	"fmt"
	"math"
	"math/bits"
)

const (
	constUint64BitCount = 64
	constUint64MaxValue = math.MaxUint64
)

// bitarray is a struct that maintains state of a bit array.
type bitarray struct {
	blks []uint64
	size int
}

// returns the index of the blk and the position of the bit within that blk
func bitPosition(index int) (int, int) {
	return index / constUint64BitCount, index % constUint64BitCount
}

// returns the number of blocks used within sieve for supplied @index
func blkCount(index int) int {
	blkIndex, bitIndex := bitPosition(index)
	if bitIndex > 0 {
		blkIndex++
	}
	return blkIndex
}

func newBitArray(size, capacity int) *bitarray {
	return &bitarray{make([]uint64, blkCount(capacity)), size}
}

// NewSieve creates a new instance of a bit array of size @size.
func NewSieve(size, capacity int) Sieve {
	return newBitArray(size, capacity)
}

func (ba *bitarray) Capacity() int {
	return len(ba.blks) * constUint64BitCount
}

func (ba *bitarray) Size() int {
	return ba.size
}

func (ba *bitarray) Resize(size int) error {
	// throw error if new size greater available capacity
	if capacity := ba.Capacity(); capacity < size {
		return ResizeOutOfRangeError{capacity, size}
	}
	// only clear active blocks
	blkCount := blkCount(size)
	for i := 0; i < blkCount; i++ {
		ba.blks[i] = 0
	}
	// resize
	ba.size = size
	return nil
}

func (ba *bitarray) Count() int {
	count := 0
	blkIndex, bitIndex := bitPosition(ba.size)
	for i := 0; i < blkIndex; i++ {
		count += bits.OnesCount64(ba.blks[i])
	}
	if bitIndex > 0 {
		blk := ba.blks[blkIndex]
		for i := 0; i < bitIndex; i++ {
			if (blk & (1 << uint(i))) != 0 {
				count++
			}
		}
	}
	return count
}

func (ba *bitarray) Set(index int) error {
	if index >= ba.size {
		return PositionOutOfRangeError{index, ba.size}
	}
	blkIndex, bitIndex := bitPosition(index)
	ba.blks[blkIndex] |= (1 << uint(bitIndex))
	return nil
}

func (ba *bitarray) Clear(index int) error {
	if index >= ba.size {
		return PositionOutOfRangeError{index, ba.size}
	}
	blkIndex, bitIndex := bitPosition(index)
	ba.blks[blkIndex] &= ^(1 << uint(bitIndex))
	return nil
}

func (ba *bitarray) Toggle(index int) error {
	if index >= ba.size {
		return PositionOutOfRangeError{index, ba.size}
	}
	blkIndex, bitIndex := bitPosition(index)
	ba.blks[blkIndex] ^= (1 << uint(bitIndex))
	return nil
}

func (ba *bitarray) SetAll() {
	blkCount := blkCount(ba.Size())
	for i := 0; i < blkCount; i++ {
		ba.blks[i] = constUint64MaxValue
	}
}

func (ba *bitarray) ClearAll() {
	blkCount := blkCount(ba.Size())
	for i := 0; i < blkCount; i++ {
		ba.blks[i] = 0
	}
}

func (ba *bitarray) ToggleAll() {
	blkCount := blkCount(ba.Size())
	for i := 0; i < blkCount; i++ {
		ba.blks[i] = ^ba.blks[i]
	}
}

func (ba *bitarray) IsSet(index int) bool {
	blkIndex, bitIndex := bitPosition(index)
	return (ba.blks[blkIndex] & (1 << uint(bitIndex))) != 0
}

func (ba *bitarray) ToNums() []int {
	blkIndex, bitIndex := bitPosition(ba.size)
	nums := make([]int, 0, (blkIndex+1)*8)
	// whole blocks
	for blkInd := 0; blkInd < blkIndex; blkInd++ {
		blk, blkPos := ba.blks[blkInd], blkInd*constUint64BitCount
		if blk == 0 {
			continue
		}
		// TODO use leading and trailing zeros in bits package
		// debugNums1, debugNums2 := make([]int, 0, 30), make([]int, 0, 30)
		bitLead, bitTail := constUint64BitCount-bits.LeadingZeros64(blk), bits.TrailingZeros64(blk)
		for bitInd := bitTail; bitInd < bitLead; bitInd++ {
			if blk&(1<<uint(bitInd)) != 0 {
				nums = append(nums, blkPos+bitInd)
				// debugNums1 = append(debugNums1, bitInd)
			}
		}
		// for bitInd := 0; bitInd < constUint64BitCount; bitInd++ {
		// 	if blk&(1<<uint(bitInd)) != 0 {
		// 		nums = append(nums, blkPos+bitInd)
		// 		// debugNums2 = append(debugNums2, bitInd)
		// 	}
		// }
		// fmt.Printf("%064b ; Tail:%d, Lead:%d, Count:%d\n", blk, bitTail, bitLead, bits.OnesCount64(blk))
		// fmt.Printf("%d %v\n", len(debugNums1), debugNums1)
		// fmt.Printf("%d %v\n\n", len(debugNums2), debugNums2)
	}
	// last part block
	if bitIndex > 0 {
		blk, blkPos := ba.blks[blkIndex], blkIndex*constUint64BitCount
		for bitInd := 0; bitInd < bitIndex; bitInd++ {
			if blk&(1<<uint(bitInd)) != 0 {
				nums = append(nums, blkPos+bitInd)
			}
		}
	}
	return nums
}

func (ba *bitarray) SubsetOf(super *bitarray) bool {
	isSubset := true
	blkCount := blkCount(ba.Size())
	for i := 0; i < blkCount && isSubset; i++ {
		isSubset = isSubset && (ba.blks[i]&super.blks[i] == ba.blks[i])
	}
	return isSubset
}

func (ba *bitarray) SetSeries(startIndex, stepIndex int) {
	for i := startIndex; i < ba.size; i += stepIndex {
		ba.Set(i)
	}
}

func (ba *bitarray) ClearSeries(startIndex, stepIndex int) {
	for i := startIndex; i < ba.size; i += stepIndex {
		ba.Clear(i)
	}
}

func (ba *bitarray) PrintRange(frIndex, toIndex int) string {
	blkString, output := fmt.Sprintf("%%0%db\n", constUint64BitCount), ""
	blkCountStt, blkCountEnd := blkCount(frIndex), blkCount(toIndex)
	for i := blkCountStt; i <= blkCountEnd; i++ {
		output += fmt.Sprintf(blkString, ba.blks[i])
	}
	return output
}
