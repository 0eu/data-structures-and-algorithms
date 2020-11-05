package dynamicarray

import "testing"

func assertCorrectCapacity(t *testing.T, array *DynamicArray, expected int32) {
	t.Helper()
	actual := array.Capacity()
	if actual != expected {
		t.Errorf("expected capacity '%d', but got '%d'", expected, actual)
	}
}

func assertCorrectLength(t *testing.T, array *DynamicArray, expected int32) {
	t.Helper()
	actual := array.Size()
	if actual != expected {
		t.Errorf("expected length '%d', but got '%d'", expected, actual)
	}
}

func assertError(t *testing.T, actual, expected error) {
	t.Helper()
	if actual != expected {
		t.Errorf("expected error '%s', but got '%s'", expected, actual)
	}
}

func TestNewDynamicArray(t *testing.T) {
	array := NewDynamicArray()
	assertCorrectLength(t, array, 0)
	assertCorrectCapacity(t, array, defaultCapacity)
}

func TestNewDynamicArrayWithCapacity(t *testing.T) {
	t.Run("initialize an array with negative capacity", func(t *testing.T) {
		_, err := NewDynamicArrayWithCapacity(-10)
		assertError(t, err, ErrorExceededCapacitySize)
	})

	t.Run("initialize an array with exceeded capacity", func(t *testing.T) {
		_, err := NewDynamicArrayWithCapacity(maxCapacity)
		assertError(t, err, ErrorExceededCapacitySize)
	})

	t.Run("initialize an array with allowed capacity", func(t *testing.T) {
		var desiredCapacity int32 = 10
		array, err := NewDynamicArrayWithCapacity(desiredCapacity)
		assertCorrectLength(t, array, int32(0))
		assertCorrectCapacity(t, array, desiredCapacity)
		assertError(t, err, nil)
	})
}
