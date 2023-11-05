package set

import (
	"fmt"
	"strings"
)

type hashSet[C comparable] map[C]struct{}

func NewHashSet[C comparable]() Set[C] {
	return hashSet[C]{}
}

func (h hashSet[C]) Add(key C) bool {
	i := len(h)
	h[key] = struct{}{}

	return i < len(h)
}

func (h hashSet[C]) AddAll(keys ...C) bool {
	i := len(h)
	for _, v := range keys {
		h[v] = struct{}{}
	}

	return i < len(h)
}

func (h hashSet[C]) Clear() {
	for k := range h {
		delete(h, k)
	}
}

func (h hashSet[C]) Contains(key C) bool {
	_, ok := h[key]
	return ok
}

func (h hashSet[C]) ContainsAll(keys ...C) bool {
	for _, v := range keys {
		if _, ok := h[v]; !ok {
			return false
		}
	}
	return true
}

func (h hashSet[C]) Equals(keys Set[C]) bool {
	second := keys.(hashSet[C])

	if h.Size() != second.Size() {
		return false
	}

	for k := range second {
		if !h.Contains(k) {
			return false
		}
	}
	return true
}

func (h hashSet[C]) IsEmpty() bool {
	return h.Size() == 0
}

func (h hashSet[C]) Remove(key C) bool {
	i := len(h)
	delete(h, key)

	return i > len(h)
}

func (h hashSet[C]) RemoveAll(keys ...C) bool {
	i := len(h)
	for _, v := range keys {
		delete(h, v)
	}

	return i > len(h)
}

func (h hashSet[C]) Size() int {
	return len(h)
}

func (h hashSet[C]) ToArray() []C {
	if h.Size() == 0 {
		return make([]C, 0)
	}

	var arr []C
	for k := range h {
		arr = append(arr, k)
	}

	return arr
}

func (h hashSet[C]) ToString() string {
	if h.Size() == 0 {
		return "Set{}"
		
	}
	setString := make([]string, 0, len(h))

	for k := range h {
		setString = append(setString, fmt.Sprintf("%v", k))
	}
	return fmt.Sprintf("Set{%s}", strings.Join(setString, ", "))
}
