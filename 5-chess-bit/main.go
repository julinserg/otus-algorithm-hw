package main

import (
	"fmt"
)

func kingMoves(pos int) uint64 {
	var k, kl, kr, nl, nr, mask uint64
	k = 1 << pos
	nr = 0x7f7f7f7f7f7f7f7f
	nl = 0xfefefefefefefefe
	kl = k & nl
	kr = k & nr
	mask = (kl << 7) | (k << 8) | (kr << 9) | (kl >> 1) | (kr << 1) | (kl >> 9) | (k >> 8) | (kr >> 7)
	return mask
}

func knightMoves(pos int) uint64 {
	var k, nA, nAB, nH, nGH, mask uint64
	k = 1 << pos
	nA = 0xFeFeFeFeFeFeFeFe
	nAB = 0xFcFcFcFcFcFcFcFc
	nH = 0x7f7f7f7f7f7f7f7f
	nGH = 0x3f3f3f3f3f3f3f3f
	mask = nGH&(k<<6|k>>10) | nH&(k<<15|k>>17) | nA&(k<<17|k>>15) | nAB&(k<<10|k>>6)
	return mask
}

func popcnt1(mask uint64) int {
	result := 0
	for mask > 0 {
		result += int(mask & 1)
		mask >>= 1
	}
	return result
}

func popcnt2(mask uint64) int {
	result := 0
	for mask > 0 {
		result++
		mask &= mask - 1
	}
	return result
}

func main() {
	fmt.Println(kingMoves(27), popcnt1(kingMoves(27)))
	fmt.Println(kingMoves(16), popcnt1(kingMoves(16)))
	fmt.Println(kingMoves(0), popcnt1(kingMoves(0)))

	fmt.Println(kingMoves(27), popcnt2(kingMoves(27)))
	fmt.Println(kingMoves(16), popcnt2(kingMoves(16)))
	fmt.Println(kingMoves(0), popcnt2(kingMoves(0)))

	fmt.Println(knightMoves(27), popcnt1(knightMoves(27)))
	fmt.Println(knightMoves(16), popcnt1(knightMoves(16)))
	fmt.Println(knightMoves(0), popcnt1(knightMoves(0)))
}
