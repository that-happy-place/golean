package collections

import (
	"testing"
)

func TestHashSet_Add(t *testing.T) {
	set := NewHashSet[int]()
	if !set.Add(1) {
		t.Error("Add should return true for a new element")
	}
	if set.Add(1) {
		t.Error("Add should return false for a duplicate element")
	}
}

func TestHashSet_AddAll(t *testing.T) {
	set := NewHashSet[int]()
	other := NewCollectionFromSlice([]int{1, 2, 3})
	if !set.AddAll(other) {
		t.Error("AddAll should return true when elements are added")
	}
	if !set.Contains(1) || !set.Contains(2) || !set.Contains(3) {
		t.Error("AddAll did not correctly add all elements")
	}
}

func TestHashSet_Clear(t *testing.T) {
	set := NewHashSet[int]()
	set.Add(1)
	set.Clear()
	if !set.IsEmpty() {
		t.Error("Clear should remove all elements")
	}
}

func TestHashSet_Contains(t *testing.T) {
	set := NewHashSet[int]()
	set.Add(1)
	if !set.Contains(1) {
		t.Error("Contains should return true for an existing element")
	}
	if set.Contains(2) {
		t.Error("Contains should return false for a non-existing element")
	}
}

func TestHashSet_Remove(t *testing.T) {
	set := NewHashSet[int]()
	set.Add(1)
	if !set.Remove(1) {
		t.Error("Remove should return true for an existing element")
	}
	if set.Remove(2) {
		t.Error("Remove should return false for a non-existing element")
	}
}

func TestHashSet_RemoveAll(t *testing.T) {
	set := NewHashSet[int]()
	set.AddAll(NewCollectionFromSlice([]int{1, 2, 3}))
	toRemove := NewCollectionFromSlice([]int{2, 3})
	if !set.RemoveAll(toRemove) {
		t.Error("RemoveAll should return true when elements are removed")
	}
	if set.Contains(2) || set.Contains(3) {
		t.Error("RemoveAll did not correctly remove all specified elements")
	}
}

func TestHashSet_RetainAll(t *testing.T) {
	set := NewHashSet[int]()
	set.AddAll(NewCollectionFromSlice([]int{1, 2, 3}))
	toRetain := NewCollectionFromSlice([]int{2, 3})
	if !set.RetainAll(toRetain) {
		t.Error("RetainAll should return true when elements are removed")
	}
	if set.Contains(1) || !set.Contains(2) || !set.Contains(3) {
		t.Error("RetainAll did not correctly retain specified elements")
	}
}

func TestHashSet_ToSlice(t *testing.T) {
	set := NewHashSet[int]()
	set.AddAll(NewCollectionFromSlice([]int{1, 2, 3}))
	slice := set.ToSlice()
	if len(slice) != 3 {
		t.Error("ToSlice did not return the correct number of elements")
	}
}
