type BrowserHistoryNode struct {
	val   string
	prev *BrowserHistoryNode 
	next *BrowserHistoryNode 
}

type BrowserHistory struct {
	head *BrowserHistoryNode 
	curr *BrowserHistoryNode 
	tail *BrowserHistoryNode 
	size  int
}


func Constructor(homepage string) BrowserHistory {
    node := &BrowserHistoryNode{
		val:  homepage,
		next: nil,
		prev: nil,
	}
	return BrowserHistory{
		head: node,
		curr: node,
		tail: node,
		size: 0,
	}
}


func (this *BrowserHistory) Visit(url string)  {
    node := &BrowserHistoryNode{
		val:  url,
		next: nil,
		prev: this.curr,
	}
	this.curr.next = node
	this.curr = node
	this.tail = node
}


func (this *BrowserHistory) Back(steps int) string {
	for i := 0; i < steps; i++ {
		if this.curr.prev == nil {
			return this.curr.val
		}
		this.curr = this.curr.prev
	}
	return this.curr.val
}


func (this *BrowserHistory) Forward(steps int) string {
	for i := 0; i < steps; i++ {
		if this.curr.next == nil {
			return this.curr.val
		}
		this.curr = this.curr.next
	}
	return this.curr.val
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */