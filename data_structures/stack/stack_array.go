package stack

type stackArray struct {
	Items []Data
	Index int
}

func NewStackArray(data ...Data) Stack {
	obj := &stackArray{}
	obj.Items = data

	return obj
}

func (s *stackArray) Push(item Data) {
	s.Resize()
	s.Items[s.Index] = item
	s.Index++
}

func (s *stackArray) Pop() Data {
	if !s.IsEmpty() {
		item := s.Items[s.Index-1]
		s.Index--
		s.Resize()
		return item
	}

	return 0 //FIXME
}

func (s *stackArray) Resize() {
	if cap(s.Items)/4 > s.Index {
		s.Items = s.Items[0 : cap(s.Items)/2]
	} else if cap(s.Items) == s.Index {
		c := make([]Data, 1+s.Index*2)
		copy(c, s.Items)
		s.Items = c
	}
}

func (s *stackArray) IsEmpty() bool {
	return s.Size() == 0
}

func (s *stackArray) Size() int {
	return s.Index
}

func (s *stackArray) Iterate() <-chan Data {
	ch := make(chan Data)
	go func() {
		for {
			if s.IsEmpty() {
				break
			}
			ch <- s.Pop()
		}
		close(ch)

	}()
	return ch
}
