package set_test

import (
	"testing"

	"github.com/sanchitdeora/data-structures-algo-go/set"
	"github.com/stretchr/testify/assert"
)

func createHashSetInt() set.Set[int]{
	return set.NewHashSet[int]()
}

func createHashSetString() set.Set[string]{
	return set.NewHashSet[string]()

}

func createHashSetFloat32() set.Set[float32]{
	return set.NewHashSet[float32]()
}

func TestAdd_HashSet(t *testing.T) {
	{
		set := createHashSetString()
		ok := set.Add("element 1")
		assert.True(t, ok)
		assert.Equal(t, 1, set.Size())

		ok = set.Add("element 2")
		assert.True(t, ok)
		assert.Equal(t, 2, set.Size())

		ok = set.Add("element 2")
		assert.False(t, ok)
		assert.Equal(t, 2, set.Size())
	}
	{
		set := createHashSetInt()
		ok := set.Add(1)
		assert.True(t, ok)
		assert.Equal(t, 1, set.Size())

		ok = set.Add(1)
		assert.False(t, ok)
		assert.Equal(t, 1, set.Size())

		ok = set.Add(2)
		assert.True(t, ok)
		assert.Equal(t, 2, set.Size())
	}
	{
		set := createHashSetFloat32()
		ok := set.Add(1)
		assert.True(t, ok)
		assert.Equal(t, 1, set.Size())

		ok = set.Add(1.5)
		assert.True(t, ok)
		assert.Equal(t, 2, set.Size())

		ok = set.Add(2.0)
		assert.True(t, ok)
		assert.Equal(t, 3, set.Size())
	}
}

func TestAddAll_HashSet(t *testing.T) {
	set := createHashSetInt()
	ok := set.AddAll(1, 2, 3)
	assert.True(t, ok)
	assert.Equal(t, 3, set.Size())

	ok = set.AddAll(5, 4, 3)
	assert.True(t, ok)
	assert.Equal(t, 5, set.Size())
}

func TestClear_HashSet(t *testing.T) {
	set := createHashSetString()
	set.Add("example")
	assert.Equal(t, 1, set.Size())

	set.Clear()
	assert.Equal(t, 0, set.Size())

	set.Add("example")
	set.Add("another one")
	assert.Equal(t, 2, set.Size())
	
	set.Clear()
	assert.Equal(t, 0, set.Size())

}

func TestContains_HashSet(t *testing.T) {
	{
		set := createHashSetFloat32()
		set.Add(3.5)
		assert.Equal(t, 1, set.Size())

		assert.False(t, set.Contains(1))
		assert.True(t, set.Contains(3.5))
		assert.True(t, set.Contains(3.50))
		assert.False(t, set.Contains(3))		
	}
	{
		set := createHashSetString()
		set.Add("test")
		assert.Equal(t, 1, set.Size())

		assert.False(t, set.Contains("example"))
		assert.True(t, set.Contains("test"))
		assert.False(t, set.Contains("Test"))
	}
}

func TestContainsAll_HashSet(t *testing.T) {
	set := createHashSetInt()
	set.Add(3)
	assert.Equal(t, 1, set.Size())

	assert.False(t, set.ContainsAll(1))
	assert.False(t, set.ContainsAll(1, 2, 3))		
	
	set.Add(2)
	assert.Equal(t, 2, set.Size())

	assert.True(t, set.ContainsAll(2, 3))

}

func TestEquals_HashSet(t *testing.T) {
	set1 := createHashSetInt()
	set2 := createHashSetInt()

	assert.True(t, set1.Equals(set2))

	set1.Add(1)
	assert.False(t, set1.Equals(set2))

	set2.Add(1)
	assert.True(t, set1.Equals(set2))

	set1.Add(2)
	set2.Add(3)
	assert.False(t, set1.Equals(set2))
}

func TestIsEmpty_HashSet(t *testing.T) {
	set := createHashSetInt()

	assert.True(t, set.IsEmpty())

	set.Add(1)
	assert.False(t, set.IsEmpty())
}

func TestRemove_HashSet(t *testing.T) {
	set := createHashSetInt()
	set.Add(1)
	assert.Equal(t, 1, set.Size())

	ok := set.Remove(1)
	assert.True(t, ok)
	assert.Equal(t, 0, set.Size())

	ok = set.Remove(1)
	assert.False(t, ok)
	assert.Equal(t, 0, set.Size())

	ok = set.Remove(3)
	assert.False(t, ok)
	assert.Equal(t, 0, set.Size())
}

func TestRemoveAll_HashSet(t *testing.T) {
	set := createHashSetInt()
	set.Add(1)
	assert.Equal(t, 1, set.Size())

	ok := set.RemoveAll(1, 2)
	assert.True(t, ok)
	assert.Equal(t, 0, set.Size())

	ok = set.RemoveAll(1, 2)
	assert.False(t, ok)
	assert.Equal(t, 0, set.Size())

	ok = set.RemoveAll(3, 4, 5)
	assert.False(t, ok)
	assert.Equal(t, 0, set.Size())
}

func TestToArray_HashSet(t *testing.T) {
	set := createHashSetInt()

	assert.Equal(t, 0, len(set.ToArray()))

	set.AddAll(1, 2, 3, 4, 5)
	assert.Equal(t, 5, len(set.ToArray()))
}

func TestString_HashSet(t *testing.T) {
	set := createHashSetInt()
	
	assert.Equal(t, "Set{}", set.ToString())
	
	set.Add(1)
	assert.Equal(t, "Set{1}", set.ToString())

	set.Add(2)
	assert.Equal(t, "Set{1, 2}", set.ToString())
}
