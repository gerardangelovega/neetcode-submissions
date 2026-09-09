type CalendarNode struct {
	Start, End int
	Left, Right *CalendarNode
}
func NewCalendarNode(start, end int) *CalendarNode {
	return &CalendarNode{
		Start: start,
		End: end,
	}	
}

type MyCalendar struct {
	root *CalendarNode
}
func Constructor() MyCalendar {
    return MyCalendar{
		root: nil,
	}
}

func (this *MyCalendar) Book(startTime int, endTime int) bool {
	if this.root == nil {
		this.root = NewCalendarNode(startTime, endTime)
		return true
	}

	var book func(*CalendarNode, int, int) bool
	book = func(root *CalendarNode, start, end int) bool {
		if end <= root.Start {
			if root.Left == nil {
				root.Left = NewCalendarNode(start, end)
				return true
			}
			return book(root.Left, start, end)
		} else if start >= root.End {
			if root.Right == nil {
				root.Right = NewCalendarNode(start, end)
				return true
			}
			return book(root.Right, start, end)
		}
		return false
	}

	return book(this.root, startTime, endTime)
}


/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */

 // We can use a segment tree to represent the range of times
 // We can build it using the lowest value in input and highest value in the input
 // We can add a booked field to indicate which times are already occupied
 // When querying, we can say a time range is bookable if none of the segments are booked
 // When we book, we should mark all of the associated segements as true
 // We are effectively performing a query + update