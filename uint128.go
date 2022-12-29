// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uint128

import "math/bits"

// uint128 represents a uint128 using two uint64s.
//
// When the methods below mention a bit number, bit 0 is the most
// significant bit (in hi) and bit 127 is the lowest (lo&1).
type Uint128 struct {
	hi uint64
	lo uint64
}

// mask6 returns a uint128 bitmask with the topmost n bits of a
// 128-bit number.
func Mask6(n int) Uint128 {
	return Uint128{^(^uint64(0) >> n), ^uint64(0) << (128 - n)}
}

// isZero reports whether u == 0.
//
// It's faster than u == (uint128{}) because the compiler (as of Go
// 1.15/1.16b1) doesn't do this trick and instead inserts a branch in
// its eq alg's generated code.
func (u Uint128) IsZero() bool { return u.hi|u.lo == 0 }

// and returns the bitwise AND of u and m (u&m).
func (u Uint128) And(m Uint128) Uint128 {
	return Uint128{u.hi & m.hi, u.lo & m.lo}
}

// xor returns the bitwise XOR of u and m (u^m).
func (u Uint128) Xor(m Uint128) Uint128 {
	return Uint128{u.hi ^ m.hi, u.lo ^ m.lo}
}

// or returns the bitwise OR of u and m (u|m).
func (u Uint128) Or(m Uint128) Uint128 {
	return Uint128{u.hi | m.hi, u.lo | m.lo}
}

// not returns the bitwise NOT of u.
func (u Uint128) Not() Uint128 {
	return Uint128{^u.hi, ^u.lo}
}

// subOne returns u - 1.
func (u Uint128) SubOne() Uint128 {
	lo, borrow := bits.Sub64(u.lo, 1, 0)
	return Uint128{u.hi - borrow, lo}
}

// addOne returns u + 1.
func (u Uint128) AddOne() Uint128 {
	lo, carry := bits.Add64(u.lo, 1, 0)
	return Uint128{u.hi + carry, lo}
}

// halves returns the two uint64 halves of the uint128.
//
// Logically, think of it as returning two uint64s.
// It only returns pointers for inlining reasons on 32-bit platforms.
func (u *Uint128) Halves() [2]*uint64 {
	return [2]*uint64{&u.hi, &u.lo}
}

// bitsSetFrom returns a copy of u with the given bit
// and all subsequent ones set.
func (u Uint128) BitsSetFrom(bit uint8) Uint128 {
	return u.Or(Mask6(int(bit)).Not())
}

// bitsClearedFrom returns a copy of u with the given bit
// and all subsequent ones cleared.
func (u Uint128) BitsClearedFrom(bit uint8) Uint128 {
	return u.And(Mask6(int(bit)))
}
