package stack_test

import (
	"testing"

	"github.com/sanchitdeora/data-structures-algo-go/stack"
	"github.com/stretchr/testify/assert"
)

func createIntStack() stack.Stack[int] {
	return stack.NewStack[int]()
}

func createFloatStack() stack.Stack[float64] {
	return stack.NewStack[float64]()
}

func createStringStack() stack.Stack[string] {
	return stack.NewStack[string]()
}

func TestIntStack_LifeCycle(t *testing.T) {
	stack := createIntStack()

	assert.True(t, stack.Empty())

	val, ok := stack.Peek()
	assert.False(t, ok)
	assert.Equal(t, 0, val)

	stack.Push(1)
	assert.Equal(t, 1, stack.Size())
	assert.False(t, stack.Empty())

	val, ok = stack.Peek()
	assert.True(t, ok)
	assert.Equal(t, 1, val)

	stack.Push(2)
	assert.Equal(t, 2, stack.Size())
	assert.Equal(t, 1, stack.Search(2))
	assert.Equal(t, -1, stack.Search(6))

	val, ok = stack.Pop()
	assert.True(t, ok)
	assert.Equal(t, 2, val)

	val, ok = stack.Pop()
	assert.True(t, ok)
	assert.Equal(t, 1, val)

	val, ok = stack.Pop()
	assert.False(t, ok)
	assert.Equal(t, 0, val)
}


func TestFloatStack_LifeCycle(t *testing.T) {
	stack := createFloatStack()

	assert.True(t, stack.Empty())

	val, ok := stack.Peek()
	assert.False(t, ok)
	assert.Equal(t, 0.0, val)

	stack.Push(1)
	assert.Equal(t, 1, stack.Size())
	assert.False(t, stack.Empty())

	val, ok = stack.Peek()
	assert.True(t, ok)
	assert.Equal(t, 1.0, val)

	stack.Push(2.1)
	assert.Equal(t, 2, stack.Size())
	assert.Equal(t, 1, stack.Search(2.1))
	assert.Equal(t, -1, stack.Search(6.9))

	val, ok = stack.Pop()
	assert.True(t, ok)
	assert.Equal(t, 2.1, val)

	val, ok = stack.Pop()
	assert.True(t, ok)
	assert.Equal(t, 1.0, val)

	val, ok = stack.Pop()
	assert.False(t, ok)
	assert.Equal(t, 0.0, val)
}


func TestStringStack_LifeCycle(t *testing.T) {
	stack := createStringStack()

	assert.True(t, stack.Empty())

	val, ok := stack.Peek()
	assert.False(t, ok)
	assert.Equal(t, "", val)

	stack.Push("element 1")
	assert.Equal(t, 1, stack.Size())
	assert.False(t, stack.Empty())

	val, ok = stack.Peek()
	assert.True(t, ok)
	assert.Equal(t, "element 1", val)

	stack.Push("element 2")
	assert.Equal(t, 2, stack.Size())
	assert.Equal(t, 1, stack.Search("element 2"))
	assert.Equal(t, -1, stack.Search("element 6"))

	val, ok = stack.Pop()
	assert.True(t, ok)
	assert.Equal(t, "element 2", val)

	val, ok = stack.Pop()
	assert.True(t, ok)
	assert.Equal(t, "element 1", val)

	val, ok = stack.Pop()
	assert.False(t, ok)
	assert.Equal(t, "", val)
}
