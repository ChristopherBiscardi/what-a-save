package main

import "testing"

var (
	crcTests = [][]byte{
		[]byte("a"),
		[]byte("bcd"),
	}
	// values pulled from rattletrap for compatibility
	//
	// stack ghci library/Rattletrap.hs
	// *Rattletrap> :set -Xoverloadedstrings
	// *Rattletrap> getCrc32 "a"
	// 797376943
	crcResults = []uint32{
		797376943,
		3125698588,
	}
)

func TestGetCRC(t *testing.T) {
	for i, v := range crcTests {
		a := getCrc(v)
		if a != crcResults[i] {
			t.Errorf("For crcTests[%d] `%v`, Expected %d, got %d", i, string(v), crcResults[i], a)
		}

	}
}
