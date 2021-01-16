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

func assertEqual(t *testing.T, actual, expected interface{}) {
	t.Helper()
	if actual != expected {
		t.Errorf("expected value '%d', but got '%d'", expected, actual)
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

		assertError(t, err, ErrorExceededCapacity)
	})

	t.Run("initialize an array with exceeded capacity", func(t *testing.T) {
		_, err := NewDynamicArrayWithCapacity(maxCapacity)

		assertError(t, err, ErrorExceededCapacity)
	})

	t.Run("initialize an array with allowed capacity", func(t *testing.T) {
		array, err := NewDynamicArrayWithCapacity(10)

		assertCorrectLength(t, array, int32(0))
		assertCorrectCapacity(t, array, 10)
		assertError(t, err, nil)
	})
}

func TestAddElementsToArray(t *testing.T) {
	t.Run("add one element to array should increase length", func(t *testing.T) {
		array := NewDynamicArray()

		err := array.Add(10)

		assertError(t, err, nil)
		assertCorrectLength(t, array, 1)

		element, err := array.Get(0)

		assertError(t, err, nil)
		assertEqual(t, element, 10)
	})

	t.Run("add 4 elements should double array's capacity", func(t *testing.T) {
		array := NewDynamicArray()

		for i := 1; i <= 4; i++ {
			_ = array.Add(i)
		}

		assertCorrectCapacity(t, array, defaultCapacity<<1)
	})

	t.Run("exceed maximum number of elements increase capacity", func(t *testing.T) {
		array := NewDynamicArray()

		for i := 1; i <= int(maxCapacity); i++ {
			_ = array.Add(i)
		}

		assertCorrectCapacity(t, array, maxCapacity)

		err := array.Add(1)
		assertError(t, err, ErrorExceededCapacity)
	})
}

func TestFind(t *testing.T) {
	t.Run("find in an empty array", func(t *testing.T) {
		array := NewDynamicArray()

		_, err := array.Find(10)

		assertError(t, err, ErrorElementNotFound)
	})

	t.Run("find not existed element's index", func(t *testing.T) {
		array := NewDynamicArray()

		_ = array.Add(11)
		_, err := array.Find(10)

		assertError(t, err, ErrorElementNotFound)
	})

	t.Run("find existed element's index", func(t *testing.T) {
		array := NewDynamicArray()

		_ = array.Add(11)
		_ = array.Add(10)
		actual, err := array.Find(10)

		assertEqual(t, actual, int32(1))
		assertError(t, err, nil)
	})
}

func TestGet(t *testing.T) {
	t.Run("get an element out of array's bounds", func(t *testing.T) {
		array := NewDynamicArray()

		_, err := array.Get(100)

		assertError(t, err, ErrorIndexOutOfRange)
	})

	t.Run("get an element by an existing index", func(t *testing.T) {
		array := NewDynamicArray()

		_ = array.Add(10)
		actual, err := array.Get(0)

		assertEqual(t, actual, 10)
		assertError(t, err, nil)
	})
}

func TestReverse(t *testing.T) {
	t.Run("reverse an empty array", func(t *testing.T) {
		array := NewDynamicArray()

		array.Reverse()

		assertCorrectLength(t, array, 0)
	})

	t.Run("get an element by an existing index", func(t *testing.T) {
		array := NewDynamicArray()

		for i := 0; i < 10; i++ {
			_ = array.Add(i)
		}
		firstElement, _ := array.Get(0)
		array.Reverse()
		lastElementAfterReverse, _ := array.Get(9)

		assertEqual(t, firstElement, lastElementAfterReverse)
	})
}

func TestSet(t *testing.T) {
	array := NewDynamicArray()

	_ = array.Add(10)
	oldValue, _ := array.Get(0)
	_ = array.Set(0, 15)
	newValue, _ := array.Get(0)
	err := array.Set(100, 1337)

	assertEqual(t, oldValue, 10)
	assertEqual(t, newValue, 15)
	assertError(t, err, ErrorIndexOutOfRange)
}

func TestRemove(t *testing.T) {
	t.Run("remove out of array's range", func(t *testing.T) {
		array := NewDynamicArray()

		err := array.RemoveAt(10)

		assertError(t, err, ErrorIndexOutOfRange)
	})

	t.Run("remove from array's beginning", func(t *testing.T) {
		array := NewDynamicArray()

		_ = array.Add(10)
		err := array.RemoveAt(0)

		assertError(t, err, nil)
		assertCorrectLength(t, array, 0)
	})

	t.Run("remove from array's middle", func(t *testing.T) {
		array := NewDynamicArray()

		_ = array.Add(10)
		_ = array.Add(10)
		_ = array.Add(10)
		err := array.RemoveAt(1)

		assertCorrectLength(t, array, 2)
		assertError(t, err, nil)
	})
}
