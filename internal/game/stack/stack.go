package stack

type Stack[T any] struct {
	items []*T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(item *T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Shift() *T {
	if s.Empty() {
		return nil
	}
	item := s.items[0]
	s.items = s.items[1:]
	return item
}

func (s *Stack[T]) Pop() *T {
	if s.Empty() {
		return nil
	}
	i := len(s.items) - 1
	item := s.items[i]
	s.items = s.items[:i]
	return item
}

func (s *Stack[T]) Empty() bool {
	return s.Len() == 0
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}

func (s *Stack[T]) Clear() {
	s.items = s.items[0:0]
}
