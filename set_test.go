package set

import (
	"testing"
)

func Test_NewNonThreadSafeSet(t *testing.T) {
	mySet := NewNonThreadSafeSet()
	mySet.Add(1)
	mySet.Add(2)
	mySet.Add(3)
	mySet.Add(4)
	if mySet.Size() != 4 {
		t.Error("NewNonThreadSafeSet: expetced 4 items")
	}
}

func Test_NewNonThreadSafeSet_DuplicateItems(t *testing.T) {
	mySet := NewNonThreadSafeSet()
	mySet.Add(1)
	mySet.Add(2)
	mySet.Add(3)
	mySet.Add(4)

	mySet.Add(1)
	mySet.Add(2)
	mySet.Add(3)
	mySet.Add(4)

	if mySet.Size() != 4 {
		t.Error("Duplicate: The set created was expected have 4 items")
	}
}

func Test_Has(t *testing.T) {
	mySet := NewNonThreadSafeSet()
	mySet.Add(1)
	mySet.Add(2)
	mySet.Add(3)
	mySet.Add(4)

	mySet.RemoveAll()

	if mySet.Size() != 0 {
		t.Error("RemoveAll: the set has been removed. So no more item should be there.")
	}
}

func Test_Remove(t *testing.T) {
	mySet := NewNonThreadSafeSet()
	mySet.Add(1)
	mySet.Add(2)
	mySet.Add(3)
	mySet.Add(4)

	count := mySet.Size()

	mySet.Remove(3)

	if count-1 != mySet.Size() {
		t.Error("The Remove method doesn't work as expetced.")
	}

	mySet.Iterate(func(item interface{}) {
		if item.(int) == 3 {
			t.Error("The removed item still inside Set.")
		}
	})
}

func Test_Iterate(t *testing.T) {
	mySet := NewNonThreadSafeSet()
	mySet.Add(1)
	mySet.Add(2)
	mySet.Add(3)
	mySet.Add(4)

	count := 0

	mySet.Iterate(func(item interface{}) {
		count += item.(int)
	})

	if count != 10 {
		t.Error("Iterate does not work as expected.")
	}
}
