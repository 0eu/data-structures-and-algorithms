package dynamicarray

import (
	"errors"
)

const (
	defaultCapacity int32 = 1 << 2
	maxCapacity     int32 = 1 << 10
)

var (
	// ErrorIndexOutOfRange will be returned if a given index is out of range a container.
	ErrorIndexOutOfRange = errors.New("an array's index should be >= 0 and < length of an array")

	// ErrorElementNotFound will be returned if a needed elements is not in a container.
	ErrorElementNotFound = errors.New("there is no a needed element in an array")

	// ErrorExceededCapacity will be returned if a given capacity is out of range.
	ErrorExceededCapacity = errors.New("capacity size should be between 1 and < 1024")
)

// Array is an ADT.
type Array interface {
	Add(int32) error
	Capacity() int32
	Find(interface{}) (int32, error)
	Get(int32) (interface{}, error)
	RemoveAt(int32) error
	Reverse()
	Set(int32, interface{}) error
	Size() int32
}

// DynamicArray is an implementation of a dynamic array using "static" arrays.
type DynamicArray struct {
	container []interface{}
	length    int32
	capacity  int32
}

// Add appends an element to the end of an array. In case, there is insufficient
// capacity it will double a capacity of an array to keep append time complexity
// constant.
func (d *DynamicArray) Add(element interface{}) error {
	if d.length+1 >= d.Capacity() {
		if err := d.resize(); err != nil {
			return err
		}
	}
	d.container[d.length] = element
	d.length++
	return nil
}

func (d *DynamicArray) resize() error {
	// If capacity is already set to maximum, we can't grow further
	if d.Capacity() == maxCapacity {
		return ErrorExceededCapacity
	}

	// If capacity is less than a half of maxCapacity value, then double it.
	// Otherwise, set to maxCapacity.
	if d.Capacity() < (maxCapacity >> 1) {
		d.capacity <<= 1
	} else {
		d.capacity = maxCapacity
	}

	var tempContainer = make([]interface{}, d.Capacity())
	copy(tempContainer, d.container)
	d.container = tempContainer
	return nil
}

// Capacity returns a number of elements in an array.
func (d *DynamicArray) Capacity() int32 {
	return d.capacity
}

// Find performs linear search on an array.
func (d *DynamicArray) Find(element interface{}) (int32, error) {
	for index, value := range d.container {
		if element == value {
			return int32(index), nil
		}
	}
	return 0, ErrorElementNotFound
}

// Get returns an element by a given index.
func (d *DynamicArray) Get(index int32) (interface{}, error) {
	if err := d.checkIndex(index); err != nil {
		return nil, err
	}
	return d.container[index], nil
}

func (d *DynamicArray) checkIndex(index int32) error {
	if index >= d.length || index < 0 {
		return ErrorIndexOutOfRange
	}
	return nil
}

// RemoveAt a given index.
func (d *DynamicArray) RemoveAt(index int32) error {
	if err := d.checkIndex(index); err != nil {
		return err
	}
	newPosition := 0
	for oldPosition := 0; oldPosition < int(d.length); oldPosition++ {
		if oldPosition == int(index) {
			continue
		}
		d.container[newPosition] = d.container[oldPosition]
		newPosition++
	}
	d.container[d.length-1] = nil
	d.length--
	return nil
}

// Reverse an array using two pointers approach.
func (d *DynamicArray) Reverse() {
	for i, j := 0, int(d.length)-1; i < j; i, j = i+1, j-1 {
		d.container[i], d.container[j] = d.container[j], d.container[i]
	}
}

// Set an element by a given index.
func (d *DynamicArray) Set(index int32, value interface{}) error {
	if err := d.checkIndex(index); err != nil {
		return err
	}
	d.container[index] = value
	return nil
}

// Size returns a number of elements in an array.
func (d *DynamicArray) Size() int32 {
	return d.length
}

// NewDynamicArray creates a new dynamic array with default capacity = 4.
func NewDynamicArray() *DynamicArray {
	return &DynamicArray{
		container: make([]interface{}, defaultCapacity),
		length:    0,
		capacity:  defaultCapacity,
	}
}

// NewDynamicArrayWithCapacity creates a new dynamic array with provided capacity
// that should not exceed MaxInt32 value, and be more than 0.
func NewDynamicArrayWithCapacity(capacity int32) (*DynamicArray, error) {
	if err := checkCapacity(capacity); err != nil {
		return nil, err
	}
	return &DynamicArray{
		container: make([]interface{}, capacity),
		length:    0,
		capacity:  capacity,
	}, nil
}

func checkCapacity(capacity int32) error {
	if capacity <= 0 || capacity >= maxCapacity {
		return ErrorExceededCapacity
	}
	return nil
}
