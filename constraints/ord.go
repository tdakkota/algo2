package constraints

type Order int

const (
	Lesser  Order = -1
	Equal         = 0
	Greater       = 1
)

type Ord[T any] func(a, b T) Order
