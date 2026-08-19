import (
    "cmp"
    "slices"
)

type Stack struct {
    items []float32
}
func NewStack() *Stack {
    return &Stack { items: make([]float32, 0) }
}
func (s *Stack) Push(val float32) {
    s.items = append(s.items, val) 
}
func (s *Stack) Pop() float32 {
    res := s.Top() 
    s.items = s.items[:s.Length()-1]
    return res
}
func (s *Stack) Top() float32 {
    return s.items[s.Length()-1]
}
func (s *Stack) Length() int {
    return len(s.items)
}

func carFleet(target int, position []int, speed []int) int {
    n := len(position)

    cars := make(map[int]int)
    for i := range n {
        cars[position[i]] = speed[i]
    }
    slices.SortFunc(position, func(a, b int) int { return cmp.Compare(b, a) })
    
    etas := make([]float32, n)
    for i := range n {
        etas[i] = float32(target - position[i]) / float32(cars[position[i]] )
    }

    stack := NewStack()
    for i := range n {
        if stack.Length() == 0 {
            stack.Push(etas[i])
            continue
        }
        if etas[i] <= stack.Top() {
            continue
        }
        stack.Push(etas[i])
    }

    return stack.Length()
}

// 1: a,_,_,b,_,_,_,_,_,_
// 2: _,_,_,a,_,b,_,_,_,_
// 3: _,_,_,_,_,_,a,b,_,_
// 4: _,_,_,_,_,_,_,_,_,ab

// 10 - 1 / 3 = 3
// 10 - 4 / 2 = 3
// 3,3

// 1: c,b,_,_,a,_,_,d,_,_,_
// 2: _,c,_,b,_,_,a,_,d,_,_
// 3: _,_,c,_,_,b,_,_,a,d,_
// 3: _,_,_,c,_,_,_,b,_,_,ad

// 10 - 4 / 2 = 3
// 10 - 1 / 2 = 4.5
// 10 - 0 / 1 = 10
// 10 - 7 / 1 = 3
// 3,4.5,5,3