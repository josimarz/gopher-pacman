package ia

type Queue[T any] []*T

func NewQueue[T any]() Queue[T] {
	return []*T{}
}

func (s *Queue[T]) Enqueue(el *T) {
	*s = append(*s, el)
}

func (s *Queue[T]) Dequeue() *T {
	if s.Empty() {
		return nil
	}
	el := (*s)[0]
	*s = (*s)[1:]
	return el
}

func (s *Queue[T]) Empty() bool {
	return s.Len() == 0
}

func (s *Queue[T]) Len() int {
	return len(*s)
}

func (s *Queue[T]) Clear() {
	*s = (*s)[0:0]
}
