package arraylist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func getArrayInt() [10]int {
	return [...]int{0, 5, 1, 6, 8, 3, 5, 9, 2, 6}
}

func TestIntArrayList_Get(t *testing.T) {
	arr := getArrayInt()
	slice := arr[:]
	c := NewIntArrayList(&slice)
	assert.Equal(t, 0, c.Get(0))
	assert.Equal(t, 6, c.Get(3))
	assert.Equal(t, 6, c.Get(9))
	assert.Equal(t, 10, c.GetLength())
	assert.Panics(t, func() { c.Get(20) })
}

func TestIntArrayList_Set(t *testing.T) {
	arr := getArrayInt()
	slice := arr[:]
	c := NewIntArrayList(&slice)
	c.Set(0, 50)
	assert.Equal(t, 50, c.Get(0))
	assert.Panics(t, func() { c.Set(20, 2) })
}

func TestIntArrayList_Swap(t *testing.T) {
	arr := getArrayInt()
	slice := arr[:]
	c := NewIntArrayList(&slice)
	c.Swap(0, 1)
	assert.Equal(t, 5, c.Get(0))
	assert.Equal(t, 0, c.Get(1))
	c.Swap(1, 9)
	assert.Equal(t, 6, c.Get(1))
	assert.Equal(t, 0, c.Get(9))
}

func getArrayFloat32() [10]float32 {
	return [...]float32{0, 5, 1, 6, 8, 3, 5, 9, 2, 6}
}

func TestFloat32ArrayList_Get(t *testing.T) {
	arr := getArrayFloat32()
	slice := arr[:]
	c := NewFloat32ArrayList(&slice)
	assert.Equal(t, float32(0), c.Get(0))
	assert.Equal(t, float32(6), c.Get(3))
	assert.Equal(t, float32(6), c.Get(9))
	assert.Equal(t, 10, c.GetLength())
	assert.Panics(t, func() { c.Get(20) })
}

func TestFloat32ArrayList_Set(t *testing.T) {
	arr := getArrayFloat32()
	slice := arr[:]
	c := NewFloat32ArrayList(&slice)
	c.Set(0, float32(50))
	assert.Equal(t, float32(50), c.Get(0))
	assert.Panics(t, func() { c.Set(20, 2) })
}

func TestFloat32ArrayList_Swap(t *testing.T) {
	arr := getArrayFloat32()
	slice := arr[:]
	c := NewFloat32ArrayList(&slice)
	c.Swap(0, 1)
	assert.Equal(t, float32(5), c.Get(0))
	assert.Equal(t, float32(0), c.Get(1))
	c.Swap(1, 9)
	assert.Equal(t, float32(6), c.Get(1))
	assert.Equal(t, float32(0), c.Get(9))
}

func getArrayFloat64() [10]float64 {
	return [...]float64{0, 5, 1, 6, 8, 3, 5, 9, 2, 6}
}

func TestFloat64ArrayList_Get(t *testing.T) {
	arr := getArrayFloat64()
	slice := arr[:]
	c := NewFloat64ArrayList(&slice)
	assert.Equal(t, float64(0), c.Get(0))
	assert.Equal(t, float64(6), c.Get(3))
	assert.Equal(t, float64(6), c.Get(9))
	assert.Equal(t, 10, c.GetLength())
	assert.Panics(t, func() { c.Get(20) })
}

func TestFloat64ArrayList_Set(t *testing.T) {
	arr := getArrayFloat64()
	slice := arr[:]
	c := NewFloat64ArrayList(&slice)
	c.Set(0, float64(50))
	assert.Equal(t, float64(50), c.Get(0))
	assert.Panics(t, func() { c.Set(20, 2) })
}

func TestFloat64ArrayList_Swap(t *testing.T) {
	arr := getArrayFloat64()
	slice := arr[:]
	c := NewFloat64ArrayList(&slice)
	c.Swap(0, 1)
	assert.Equal(t, float64(5), c.Get(0))
	assert.Equal(t, float64(0), c.Get(1))
	c.Swap(1, 9)
	assert.Equal(t, float64(6), c.Get(1))
	assert.Equal(t, float64(0), c.Get(9))
}

func getArrayString() [10]string {
	return [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
}

func TestStringArrayList_Get(t *testing.T) {
	arr := getArrayString()
	slice := arr[:]
	c := NewStringArrayList(&slice)
	assert.Equal(t, "a", c.Get(0))
	assert.Equal(t, "d", c.Get(3))
	assert.Equal(t, "j", c.Get(9))
	assert.Equal(t, 10, c.GetLength())
	assert.Panics(t, func() { c.Get(20) })
}

func TestStringArrayList_Set(t *testing.T) {
	arr := getArrayString()
	slice := arr[:]
	c := NewStringArrayList(&slice)
	c.Set(0, "x")
	assert.Equal(t, "x", c.Get(0))
	assert.Panics(t, func() { c.Set(20, "y") })
}

func TestStringArrayList_Swap(t *testing.T) {
	arr := getArrayString()
	slice := arr[:]
	c := NewStringArrayList(&slice)
	c.Swap(0, 1)
	assert.Equal(t, "b", c.Get(0))
	assert.Equal(t, "a", c.Get(1))
	c.Swap(1, 9)
	assert.Equal(t, "j", c.Get(1))
	assert.Equal(t, "a", c.Get(9))
}
