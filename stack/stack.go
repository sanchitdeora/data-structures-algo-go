package stack

type Stack[C comparable] []C

func NewStack[C comparable]() Stack[C] {
	return make([]C, 0)
}

func (s *Stack[C]) Empty() bool {
	return s.Size() == 0
}

func (s *Stack[C]) Peek() (C, bool) {
	if s.Empty() {
		var result C
		return result, false
	}
	return (*s)[len(*s)-1], true
}

func (s *Stack[C]) Pop() (C, bool) {
	if s.Empty() {
		var result C
		return result, false
	}
	lastElem := (*s)[len(*s)-1]
	*s = append((*s)[:len(*s)-1])
	return lastElem, true
}

func (s *Stack[C]) Push(key C) {
	*s = append(*s, key)
}

func (s *Stack[C]) Search(key C) int {
	for i, v := range *s {
		if v == key {
			return i
		}
	}
	return -1
}

func (s Stack[C]) Size() int {
	 return len(s) 
}
