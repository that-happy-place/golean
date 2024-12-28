package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollection_Add(t *testing.T) {
	c := NewCollectionFromSlice([]int{})
	assert.True(t, c.Add(10))
	assert.Equal(t, []int{10}, c.ToSlice())
}

func TestCollection_AddAll(t *testing.T) {
	c := NewCollectionFromSlice([]int{})
	other := NewCollectionFromSlice([]int{1, 2, 3})
	assert.True(t, c.AddAll(other))
	assert.Equal(t, []int{1, 2, 3}, c.ToSlice())
}

func TestCollection_Clear(t *testing.T) {
	c := NewCollectionFromSlice([]int{1, 2, 3})
	c.Clear()
	assert.True(t, c.IsEmpty())
}

func TestCollection_Contains(t *testing.T) {
	c := NewCollectionFromSlice([]int{1, 2, 3})
	assert.True(t, c.Contains(2))
	assert.False(t, c.Contains(4))
}

func TestCollection_ContainsAll(t *testing.T) {
	c := NewCollectionFromSlice([]int{1, 2, 3})
	other := NewCollectionFromSlice([]int{1, 2})
	assert.True(t, c.ContainsAll(other))
	other = NewCollectionFromSlice([]int{4})
	assert.False(t, c.ContainsAll(other))
}

func TestCollection_Remove(t *testing.T) {
	c := NewCollectionFromSlice([]int{1, 2, 3})
	assert.True(t, c.Remove(2))
	assert.Equal(t, []int{1, 3}, c.ToSlice())
	assert.False(t, c.Remove(4))
}

func TestCollection_RemoveAll(t *testing.T) {
	c := NewCollectionFromSlice([]int{1, 2, 3})
	other := NewCollectionFromSlice([]int{1, 3})
	assert.True(t, c.RemoveAll(other))
	assert.Equal(t, []int{2}, c.ToSlice())
}

func TestCollection_RemoveIf(t *testing.T) {
	c := NewCollectionFromSlice([]int{1, 2, 3, 4})
	assert.True(t, c.RemoveIf(func(val int) bool { return val%2 == 0 }))
	assert.Equal(t, []int{1, 3}, c.ToSlice())
}

func TestCollection_Len(t *testing.T) {
	c := NewCollectionFromSlice([]int{1, 2, 3})
	assert.Equal(t, 3, c.Len())
}

func TestCollection_ToSlice(t *testing.T) {
	c := NewCollectionFromSlice([]int{1, 2, 3})
	assert.Equal(t, []int{1, 2, 3}, c.ToSlice())
}
