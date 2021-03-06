// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package gopp

import "math/rand"

func ShuffleBool(slc []bool) {
	n := len(slc)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slc[i], slc[j] = slc[j], slc[i]
	}
}

func ShuffleByte(slc []byte) {
	n := len(slc)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slc[i], slc[j] = slc[j], slc[i]
	}
}

func ShuffleComplex128(slc []complex128) {
	n := len(slc)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slc[i], slc[j] = slc[j], slc[i]
	}
}

func ShuffleComplex64(slc []complex64) {
	n := len(slc)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slc[i], slc[j] = slc[j], slc[i]
	}
}

func ShuffleError(slc []error) {
	n := len(slc)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slc[i], slc[j] = slc[j], slc[i]
	}
}

func ShuffleFloat32(slc []float32) {
	n := len(slc)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slc[i], slc[j] = slc[j], slc[i]
	}
}

func ShuffleFloat64(slc []float64) {
	n := len(slc)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slc[i], slc[j] = slc[j], slc[i]
	}
}

func ShuffleInt(slc []int) {
	n := len(slc)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slc[i], slc[j] = slc[j], slc[i]
	}
}

func ShuffleInt16(slc []int16) {
	n := len(slc)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slc[i], slc[j] = slc[j], slc[i]
	}
}

func ShuffleInt32(slc []int32) {
	n := len(slc)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slc[i], slc[j] = slc[j], slc[i]
	}
}

func ShuffleInt64(slc []int64) {
	n := len(slc)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slc[i], slc[j] = slc[j], slc[i]
	}
}

func ShuffleInt8(slc []int8) {
	n := len(slc)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slc[i], slc[j] = slc[j], slc[i]
	}
}

func ShuffleRune(slc []rune) {
	n := len(slc)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slc[i], slc[j] = slc[j], slc[i]
	}
}

func ShuffleString(slc []string) {
	n := len(slc)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slc[i], slc[j] = slc[j], slc[i]
	}
}

func ShuffleUint(slc []uint) {
	n := len(slc)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slc[i], slc[j] = slc[j], slc[i]
	}
}

func ShuffleUint16(slc []uint16) {
	n := len(slc)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slc[i], slc[j] = slc[j], slc[i]
	}
}

func ShuffleUint32(slc []uint32) {
	n := len(slc)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slc[i], slc[j] = slc[j], slc[i]
	}
}

func ShuffleUint64(slc []uint64) {
	n := len(slc)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slc[i], slc[j] = slc[j], slc[i]
	}
}

func ShuffleUint8(slc []uint8) {
	n := len(slc)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slc[i], slc[j] = slc[j], slc[i]
	}
}

func ShuffleUintptr(slc []uintptr) {
	n := len(slc)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slc[i], slc[j] = slc[j], slc[i]
	}
}
