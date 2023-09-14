package ia

type Stack[T any] []*T

func NewStack[T any]() Stack[T] {
	return []*T{}
}

func (s *Stack[T]) Push(el *T) {
	*s = append(*s, el)
}

func (s *Stack[T]) Shift() *T {
	if s.Empty() {
		return nil
	}
	el := (*s)[0]
	*s = (*s)[1:]
	return el
}

func (s *Stack[T]) Pop() *T {
	if s.Empty() {
		return nil
	}
	i := s.Len() - 1
	el := (*s)[i]
	*s = (*s)[:i]
	return el
}

func (s *Stack[T]) Empty() bool {
	return s.Len() == 0
}

func (s *Stack[T]) Len() int {
	return len(*s)
}

func (s *Stack[T]) Clear() {
	*s = (*s)[0:0]
}
