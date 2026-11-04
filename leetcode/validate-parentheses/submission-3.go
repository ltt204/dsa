type Stack struct {
	items []rune 
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	
	index := len(s.items) - 1
	val := s.items[index]
	s.items = s.items[:index]

	return val, true
}

func (s *Stack) Push(item rune) {
	s.items = append(s.items, item)
}

func isValid(s string) bool {
	stack := &Stack{}

	closeToOpen := map[rune]rune{')':'(', '}':'{', ']':'['}
	
	for _, c := range s {
		if open, ok := closeToOpen[c]; ok {
			if !stack.IsEmpty() {
				top, ok := stack.Pop()

				if ok && top != open {
					return false
				}
			} else {
				return false
			}
		} else {
			stack.Push(c)
		}
	}

	return stack.IsEmpty()
}
