package set

type Set[C comparable] interface {
	Add(key C) bool
	AddAll(keys ...C) bool
	Clear()
	Contains(key C) bool
	ContainsAll(keys ...C) bool
	Equals(key Set[C]) bool
	IsEmpty() bool
	Remove(key C) bool
	RemoveAll(keys ...C) bool
	Size() int
	ToArray() []C
	ToString() string
}