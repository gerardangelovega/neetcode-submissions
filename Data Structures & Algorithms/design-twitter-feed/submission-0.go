// Tweets should be a struct poitner
// Tweets should have a special identifier indicating recentness (auto-incrementing)
// Tweets can be stored in a Max Heap with the aforementioned special identifier as the value being
// sorted. This will allows as to Pop the latest tweets for the news feed
// User is a struct with an array of tweets they posted (possibly a heap to sort by recency)
// User is a struct with an array of users they follow 
//		- we could use hashmaps for fast access and convenient deletion
//		- we could use BST maps for fast traversal
// User is a struct with a unique id
// News feed might be an algorithm to iterate through all followed user's and check the top of each 
// followed user's heap of tweets and checks which heap has the most recent post and pop from that 
// heap
// Twitter might have an array of users or rather a map of users (hashmaps will be easier for this 
// case)
// Twitter will dynamically create users if a user does not exist yet
type Heap struct {
	items []*Tweet
	cmp func(*Tweet, *Tweet) bool
}
func NewHeap(cmp func(*Tweet, *Tweet) bool) *Heap {
	return &Heap{
		items: make([]*Tweet, 1),
		cmp: cmp,
	}
}
func (h *Heap) Push(tweet *Tweet) {
	h.items = append(h.items, tweet)
	if h.Size() == 1 {
		return
	}
	h.SiftUp(h.Size())
}
func (h *Heap) Pop() *Tweet {
	if h.Size() == 0 {
		return nil
	}
	if h.Size() == 1 {
		rv := h.items[h.Size()]
		h.items = h.items[:h.Size()]
		return rv
	}
	rv := h.items[1]
	h.items[1] = h.items[h.Size()]
	h.items = h.items[:h.Size()]
	h.SiftDown(1)
	return rv
}
func (h *Heap) Top() *Tweet {
	if h.Size() == 0 {
		return nil
	}
	return h.items[1]
}
func (h *Heap) SiftUp(i int) {
	for i > 1 && h.cmp(h.items[i], h.items[i >> 1]) {
		h.items[i], h.items[i >> 1] = h.items[i >> 1], h.items[i]
		i = i >> 1
	}
}
func (h *Heap) SiftDown(i int) {
	n := h.Size()+1

	for (i << 1) < n {
		l := (i << 1)
		r := (i << 1) | 1
		s :=  i

		if l < n && h.cmp(h.items[l], h.items[s]) {
			s = l	
		}
		if r < n && h.cmp(h.items[r], h.items[s]) {
			s = r	
		}
		if s == i {
			break
		}
		h.items[i], h.items[s] = h.items[s], h.items[i]
		i = s
	}
}
func (h *Heap) Size() int {
	return len(h.items)-1
}

type User struct {
	id int
	tweets map[int]*Tweet
	following map[int]*User
}
func NewUser(id int) *User {
	return &User{
		id: id,
		tweets: make(map[int]*Tweet),
		following: make(map[int]*User),
	}
}
func (u *User) Tweet(tweet *Tweet) {
	u.tweets[tweet.id] = tweet
}
func (u *User) Follow(user *User) {
	u.following[user.id] = user
}
func (u *User) Unfollow(user *User) {
	delete(u.following, user.id)	
}

type Tweet struct {
	id int
	recency int	
}
func NewTweet(id, recency int) *Tweet {
	return &Tweet{
		id: id,
		recency: recency,
	}
}


type Twitter struct {
    users map[int]*User
	recencyCounter int
}
func Constructor() Twitter {
    return Twitter{
		users: make(map[int]*User, 0),
		recencyCounter: 0,
	}
}
func (this *Twitter) PostTweet(userId int, tweetId int)  {
	var user *User
	var exists bool

	if user, exists = this.users[userId]; !exists {
		this.users[userId] = NewUser(userId)
		user = this.users[userId]
	}
	user.Tweet(NewTweet(tweetId, this.getRecency()))
}
func (this *Twitter) GetNewsFeed(userId int) []int {
	var user *User
	var exists bool

    heap := NewHeap(func (a, b *Tweet) bool {
		return a.recency < b.recency
	})

	if user, exists = this.users[userId]; !exists {
		return []int{}
	}

	for _, followee := range user.following {
		for _, tweet := range followee.tweets {
			heap.Push(tweet)

			if heap.Size() > 10 {
				heap.Pop()
			}
		}
	}
	for _, tweet := range user.tweets {
		heap.Push(tweet)

		if heap.Size() > 10 {
			heap.Pop()
		}
	}

	n := heap.Size()
	res := make([]int, n)

	for i := n-1; i > -1; i-- {
		res[i] = heap.Pop().id
	}

	return res
}
func (this *Twitter) Follow(followerId int, followeeId int)  {
	if followerId == followeeId {
		return
	}
	var follower *User
	var followee *User
	var exists bool

	if follower, exists = this.users[followerId]; !exists {
		this.users[followerId] = NewUser(followerId)
		follower = this.users[followerId]
	}
	if followee, exists = this.users[followeeId]; !exists {
		this.users[followeeId] = NewUser(followeeId)
		followee = this.users[followeeId]
	}
	follower.Follow(followee)
}
func (this *Twitter) Unfollow(followerId int, followeeId int)  {
	if followerId == followeeId {
		return
	}
	var follower *User
	var followee *User
	var exists bool

	if follower, exists = this.users[followerId]; !exists {
		return
	}
	if followee, exists = this.users[followeeId]; !exists {
		return
	}
	follower.Unfollow(followee)
}
func (this *Twitter) getRecency() int {
	rv := this.recencyCounter
	this.recencyCounter++
	return rv
}
