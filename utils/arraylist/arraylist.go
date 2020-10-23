package arraylist

type ArrayList interface {
	Get(index int) interface{}
	Set(index int, value interface{})
	Swap(from int, to int)
	GetLength() int
}

type AbstractArrayList struct {
	ArrayList
}

func (list AbstractArrayList) Swap(from int, to int) {
	v := list.Get(to)
	list.Set(to, list.Get(from))
	list.Set(from, v)
}

type IntArrayList struct {
	data []int
	AbstractArrayList
}

func NewIntArrayList(arr []int) IntArrayList {
	list := IntArrayList{arr, AbstractArrayList{}}
	list.AbstractArrayList.ArrayList = list
	return list
}

func (list IntArrayList) Get(index int) interface{} {
	return list.data[index]
}

func (list IntArrayList) Set(index int, value interface{}) {
	list.data[index] = value.(int)
}

func (list IntArrayList) GetLength() int {
	return len(list.data)
}

type Float32ArrayList struct {
	data []float32
	AbstractArrayList
}

func NewFloat32ArrayList(arr []float32) Float32ArrayList {
	list := Float32ArrayList{arr, AbstractArrayList{}}
	list.AbstractArrayList.ArrayList = list
	return list
}

func (list Float32ArrayList) Get(index int) interface{} {
	return list.data[index]
}

func (list Float32ArrayList) Set(index int, value interface{}) {
	list.data[index] = value.(float32)
}

func (list Float32ArrayList) GetLength() int {
	return len(list.data)
}

type Float64ArrayList struct {
	data []float64
	AbstractArrayList
}

func NewFloat64ArrayList(arr []float64) Float64ArrayList {
	list := Float64ArrayList{arr, AbstractArrayList{}}
	list.AbstractArrayList.ArrayList = list
	return list
}

func (list Float64ArrayList) Get(index int) interface{} {
	return list.data[index]
}

func (list Float64ArrayList) Set(index int, value interface{}) {
	list.data[index] = value.(float64)
}

func (list Float64ArrayList) GetLength() int {
	return len(list.data)
}

type StringArrayList struct {
	data []string
	AbstractArrayList
}

func NewStringArrayList(arr []string) StringArrayList {
	list := StringArrayList{arr, AbstractArrayList{}}
	list.AbstractArrayList.ArrayList = list
	return list
}

func (list StringArrayList) Get(index int) interface{} {
	return list.data[index]
}

func (list StringArrayList) Set(index int, value interface{}) {
	list.data[index] = value.(string)
}

func (list StringArrayList) GetLength() int {
	return len(list.data)
}
