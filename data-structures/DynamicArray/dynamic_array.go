package dynamicarray

import (
	"errors"
	"math"
)

const (
	defaultCapacity int32 = 1 << 2
	maxCapacity     int32 = math.MaxInt32
)

var (
	// ErrorIndexOutOfRange will be returned if a given index is out of range a container.
	ErrorIndexOutOfRange = errors.New("An array's index should be >= 0 and < length of an array")

	// ErrorElementIsNotInArray will be returned if a needed elements is not in a container.
	ErrorElementIsNotInArray = errors.New("There is no a needed element in an array")

	// ErrorExceededCapacitySize will be returned if a given capacity is out of range.
	ErrorExceededCapacitySize = errors.New("Capacity size should be between and < 2147483647")
)

// Array is an
type Array interface {
	Add(int32) error
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
	if d.length+1 >= d.capacity {
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
	if d.capacity == maxCapacity {
		return ErrorExceededCapacitySize
	}

	// If capacity is less than a half of maxCapacity value, then double it.
	// Otherwise, set to maxCapacity.
	if d.capacity < (maxCapacity >> 1) {
		d.capacity <<= 1
	} else {
		d.capacity = maxCapacity
	}

	var tempContainer = make([]interface{}, d.capacity)
	copy(tempContainer, d.container)
	d.container = tempContainer
	return nil
}

// Find performs linear search on an array.
func (d *DynamicArray) Find(element interface{}) (int32, error) {
	for index, value := range d.container {
		if element == value {
			return int32(index), nil
		}
	}
	return 0, ErrorElementIsNotInArray
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
	for i, j := 0, int(d.length); i < j; i, j = i+1, j-1 {
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
func NewDynamicArray() (*DynamicArray, error) {
	return &DynamicArray{
		container: make([]interface{}, defaultCapacity),
		length:    0,
		capacity:  defaultCapacity,
	}, nil
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
		return ErrorExceededCapacitySize
	}
	return nil
}
