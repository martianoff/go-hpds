package comparator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const fl32_1 float32 = 1
const fl32_2 float32 = 2
const fl32_3 float32 = 1
const fl32_4 float32 = 3
const fl32_5 float64 = 3

func TestFloat32Comparator_Equal(t *testing.T) {
	c := NewFloat32Comparator()
	assert.Equal(t, true, c.Equal(fl32_1, fl32_3))
	assert.Equal(t, false, c.Equal(fl32_1, fl32_2))
	assert.Equal(t, false, c.Equal(fl32_2, fl32_4))
	assert.Panics(t, func() { c.Equal(fl32_4, fl32_5) })
}

func TestFloat32Comparator_Greater(t *testing.T) {
	c := NewFloat32Comparator()
	assert.Equal(t, false, c.Greater(fl32_1, fl32_3))
	assert.Equal(t, true, c.Greater(fl32_2, fl32_1))
	assert.Equal(t, false, c.Greater(fl32_2, fl32_4))
	assert.Panics(t, func() { c.Greater(fl32_4, fl32_5) })
}

func TestFloat32Comparator_GreaterOrEqual(t *testing.T) {
	c := NewFloat32Comparator()
	assert.Equal(t, true, c.GreaterOrEqual(fl32_1, fl32_3))
	assert.Equal(t, false, c.GreaterOrEqual(fl32_1, fl32_2))
	assert.Equal(t, true, c.GreaterOrEqual(fl32_4, fl32_2))
	assert.Panics(t, func() { c.GreaterOrEqual(fl32_4, fl32_5) })
}

func TestFloat32Comparator_Less(t *testing.T) {
	c := NewFloat32Comparator()
	assert.Equal(t, false, c.Less(fl32_1, fl32_3))
	assert.Equal(t, true, c.Less(fl32_1, fl32_2))
	assert.Equal(t, true, c.Less(fl32_2, fl32_4))
	assert.Panics(t, func() { c.Less(fl32_4, fl32_5) })
}

func TestFloat32Comparator_LessOrEqual(t *testing.T) {
	c := NewFloat32Comparator()
	assert.Equal(t, true, c.LessOrEqual(fl32_1, fl32_3))
	assert.Equal(t, false, c.LessOrEqual(fl32_2, fl32_1))
	assert.Equal(t, true, c.LessOrEqual(fl32_2, fl32_4))
	assert.Panics(t, func() { c.LessOrEqual(fl32_4, fl32_5) })
}

const fl64_1 float64 = 1
const fl64_2 float64 = 2
const fl64_3 float64 = 1
const fl64_4 float64 = 3
const fl64_5 float32 = 3

func TestFloat64Comparator_Equal(t *testing.T) {
	c := NewFloat64Comparator()
	assert.Equal(t, true, c.Equal(fl64_1, fl64_3))
	assert.Equal(t, false, c.Equal(fl64_1, fl64_2))
	assert.Equal(t, false, c.Equal(fl64_2, fl64_4))
	assert.Panics(t, func() { c.Equal(fl64_4, fl64_5) })
}

func TestFloat64Comparator_Greater(t *testing.T) {
	c := NewFloat64Comparator()
	assert.Equal(t, false, c.Greater(fl64_1, fl64_3))
	assert.Equal(t, true, c.Greater(fl64_2, fl64_1))
	assert.Equal(t, false, c.Greater(fl64_2, fl64_4))
	assert.Panics(t, func() { c.Greater(fl64_4, fl64_5) })
}

func TestFloat64Comparator_GreaterOrEqual(t *testing.T) {
	c := NewFloat64Comparator()
	assert.Equal(t, true, c.GreaterOrEqual(fl64_1, fl64_3))
	assert.Equal(t, false, c.GreaterOrEqual(fl64_1, fl64_2))
	assert.Equal(t, true, c.GreaterOrEqual(fl64_4, fl64_2))
	assert.Panics(t, func() { c.GreaterOrEqual(fl64_4, fl64_5) })
}

func TestFloat64Comparator_Less(t *testing.T) {
	c := NewFloat64Comparator()
	assert.Equal(t, false, c.Less(fl64_1, fl64_3))
	assert.Equal(t, true, c.Less(fl64_1, fl64_2))
	assert.Equal(t, true, c.Less(fl64_2, fl64_4))
	assert.Panics(t, func() { c.Less(fl64_4, fl64_5) })
}

func TestFloat64Comparator_LessOrEqual(t *testing.T) {
	c := NewFloat64Comparator()
	assert.Equal(t, true, c.LessOrEqual(fl64_1, fl64_3))
	assert.Equal(t, false, c.LessOrEqual(fl64_2, fl64_1))
	assert.Equal(t, true, c.LessOrEqual(fl64_2, fl64_4))
	assert.Panics(t, func() { c.LessOrEqual(fl64_4, fl64_5) })
}

const int1 int = 1
const int2 int = 2
const int3 int = 1
const int4 int = 3
const int5 float32 = 3

func TestIntComparator_Equal(t *testing.T) {
	c := NewIntComparator()
	assert.Equal(t, true, c.Equal(int1, int3))
	assert.Equal(t, false, c.Equal(int1, int2))
	assert.Equal(t, false, c.Equal(int2, int4))
	assert.Panics(t, func() { c.Equal(int4, int5) })
}

func TestIntComparator_Greater(t *testing.T) {
	c := NewIntComparator()
	assert.Equal(t, false, c.Greater(int1, int3))
	assert.Equal(t, true, c.Greater(int2, int1))
	assert.Equal(t, false, c.Greater(int2, int4))
	assert.Panics(t, func() { c.Greater(int4, int5) })
}

func TestIntComparator_GreaterOrEqual(t *testing.T) {
	c := NewIntComparator()
	assert.Equal(t, true, c.GreaterOrEqual(int1, int3))
	assert.Equal(t, false, c.GreaterOrEqual(int1, int2))
	assert.Equal(t, true, c.GreaterOrEqual(int4, int2))
	assert.Panics(t, func() { c.GreaterOrEqual(int4, int5) })
}

func TestIntComparator_Less(t *testing.T) {
	c := NewIntComparator()
	assert.Equal(t, false, c.Less(int1, int3))
	assert.Equal(t, true, c.Less(int1, int2))
	assert.Equal(t, true, c.Less(int2, int4))
	assert.Panics(t, func() { c.Less(int4, int5) })
}

func TestIntComparator_LessOrEqual(t *testing.T) {
	c := NewIntComparator()
	assert.Equal(t, true, c.LessOrEqual(int1, int3))
	assert.Equal(t, false, c.LessOrEqual(int2, int1))
	assert.Equal(t, true, c.LessOrEqual(int2, int4))
	assert.Panics(t, func() { c.LessOrEqual(int4, int5) })
}

const string1 string = "a"
const string2 string = "b"
const string3 string = "a"
const string4 string = "c"
const string5 int = 3

func TestStringComparator_Equal(t *testing.T) {
	c := NewStringComparator()
	assert.Equal(t, true, c.Equal(string1, string3))
	assert.Equal(t, false, c.Equal(string1, string2))
	assert.Equal(t, false, c.Equal(string2, string4))
	assert.Panics(t, func() { c.Equal(string4, string5) })
}

func TestStringComparator_Greater(t *testing.T) {
	c := NewStringComparator()
	assert.Equal(t, false, c.Greater(string1, string3))
	assert.Equal(t, true, c.Greater(string2, string1))
	assert.Equal(t, false, c.Greater(string2, string4))
	assert.Panics(t, func() { c.Greater(string4, string5) })
}

func TestStringComparator_GreaterOrEqual(t *testing.T) {
	c := NewStringComparator()
	assert.Equal(t, true, c.GreaterOrEqual(string1, string3))
	assert.Equal(t, false, c.GreaterOrEqual(string1, string2))
	assert.Equal(t, true, c.GreaterOrEqual(string4, string2))
	assert.Panics(t, func() { c.GreaterOrEqual(string4, string5) })
}

func TestStringComparator_Less(t *testing.T) {
	c := NewStringComparator()
	assert.Equal(t, false, c.Less(string1, string3))
	assert.Equal(t, true, c.Less(string1, string2))
	assert.Equal(t, true, c.Less(string2, string4))
	assert.Panics(t, func() { c.Less(string4, string5) })
}

func TestStringComparator_LessOrEqual(t *testing.T) {
	c := NewStringComparator()
	assert.Equal(t, true, c.LessOrEqual(string1, string3))
	assert.Equal(t, false, c.LessOrEqual(string2, string1))
	assert.Equal(t, true, c.LessOrEqual(string2, string4))
	assert.Panics(t, func() { c.LessOrEqual(string4, string5) })
}
