package constraints

type Ordered interface {
	type int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, string
}

// Real permits any integer or float type.
type Real interface {
	type int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64
}

// Complex permits any complex type.
type Complex interface {
	type complex64, complex128
}

// Float permits any float type.
type Float interface {
	type float32, float64
}

// Integer permits any integer type.
type Integer interface {
	type int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr
}

// Signed permits any signed integer type.
type Signed interface {
	type int, int8, int16, int32, int64
}

// Unsigned permits any unsigned integer type.
type Unsigned interface {
	type uint, uint8, uint16, uint32, uint64, uintptr
}

type IntIndex[T any] interface {
	type []T, map[int]T
}