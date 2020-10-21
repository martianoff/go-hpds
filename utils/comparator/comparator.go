package comparator

type Comparator interface {
	Less(value1 interface{}, value2 interface{}) bool
	Greater(value1 interface{}, value2 interface{}) bool
	LessOrEqual(value1 interface{}, value2 interface{}) bool
	GreaterOrEqual(value1 interface{}, value2 interface{}) bool
	Equal(value1 interface{}, value2 interface{}) bool
}

type AbstractComparator struct{ Comparator }

func (cmp AbstractComparator) Greater(value1 interface{}, value2 interface{}) bool {
	return !cmp.Less(value1, value2) && !cmp.Equal(value1, value2)
}

func (cmp AbstractComparator) LessOrEqual(value1 interface{}, value2 interface{}) bool {
	return cmp.Less(value1, value2) || cmp.Equal(value1, value2)
}

func (cmp AbstractComparator) GreaterOrEqual(value1 interface{}, value2 interface{}) bool {
	return !cmp.Less(value1, value2)
}

type IntComparator struct{ AbstractComparator }

func NewIntComparator() *IntComparator {
	cmp := IntComparator{AbstractComparator{}}
	cmp.AbstractComparator.Comparator = cmp
	return &cmp
}

func (cmp IntComparator) Less(value1 interface{}, value2 interface{}) bool {
	return value1.(int) < value2.(int)
}

func (cmp IntComparator) Equal(value1 interface{}, value2 interface{}) bool {
	return value1.(int) == value2.(int)
}

type Float32Comparator struct{ AbstractComparator }

func NewFloat32Comparator() *Float32Comparator {
	cmp := Float32Comparator{AbstractComparator{}}
	cmp.AbstractComparator.Comparator = cmp
	return &cmp
}

func (cmp Float32Comparator) Less(value1 interface{}, value2 interface{}) bool {
	return value1.(float32) < value2.(float32)
}

func (cmp Float32Comparator) Equal(value1 interface{}, value2 interface{}) bool {
	return value1.(float32) == value2.(float32)
}

type Float64Comparator struct{ AbstractComparator }

func NewFloat64Comparator() *Float64Comparator {
	cmp := Float64Comparator{AbstractComparator{}}
	cmp.AbstractComparator.Comparator = cmp
	return &cmp
}

func (cmp Float64Comparator) Less(value1 interface{}, value2 interface{}) bool {
	return value1.(float64) < value2.(float64)
}

func (cmp Float64Comparator) Equal(value1 interface{}, value2 interface{}) bool {
	return value1.(float64) == value2.(float64)
}

type StringComparator struct{ AbstractComparator }

func NewStringComparator() *StringComparator {
	cmp := StringComparator{AbstractComparator{}}
	cmp.AbstractComparator.Comparator = cmp
	return &cmp
}

func (cmp StringComparator) Less(value1 interface{}, value2 interface{}) bool {
	return value1.(string) < value2.(string)
}

func (cmp StringComparator) Equal(value1 interface{}, value2 interface{}) bool {
	return value1.(string) == value2.(string)
}
