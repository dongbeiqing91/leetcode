package designtwitter

// https://leetcode.cn/problems/design-twitter/

import "container/heap"

type Twitter struct {
	GlobalTimestamp int
	IdUserMapping   map[int]*User
}

type Tweet struct {
	Id        int
	UserId    int
	Timestamp int
	Next      *Tweet
}

func Constructor() Twitter {
	return Twitter{
		IdUserMapping: make(map[int]*User),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.GlobalTimestamp++
	timestamp := this.GlobalTimestamp
	var user *User
	u, ok := this.IdUserMapping[userId]
	if !ok {
		user = &User{Id: userId}
		this.IdUserMapping[userId] = user
	} else {
		user = u
	}
	t := &Tweet{Id: tweetId, UserId: userId, Timestamp: timestamp}
	if user.TweetsHead == nil {
		user.TweetsHead = t
	} else {
		t.Next = user.TweetsHead
		user.TweetsHead = t
	}
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	user, ok := this.IdUserMapping[userId]
	if !ok {
		return []int{}
	}
	users := []*User{user}
	for _, v := range user.Followings {
		users = append(users, v)
	}
	pq := &PriorityQueue{}
	heap.Init(pq)
	for _, u := range users {
		if u.TweetsHead != nil {
			heap.Push(pq, u.TweetsHead)
		}
	}
	k := 10
	var res []int
	for pq.Len() > 0 && k > 0 {
		t := heap.Pop(pq).(*Tweet)
		res = append(res, t.Id)
		if t.Next != nil {
			heap.Push(pq, t.Next)
		}
		k--
	}
	return res
}

var _ heap.Interface = &PriorityQueue{}

type PriorityQueue []*Tweet

func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) Less(i, j int) bool {
	return p[i].Timestamp > p[j].Timestamp
}

func (p *PriorityQueue) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *PriorityQueue) Push(x any) {
	*p = append(*p, x.(*Tweet))
}

func (p *PriorityQueue) Pop() any {
	pLen := p.Len()
	res := (*p)[pLen-1]
	*p = (*p)[0 : pLen-1]
	return res
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	if _, ok := this.IdUserMapping[followerId]; !ok {
		this.IdUserMapping[followerId] = &User{
			Id: followerId,
		}
	}
	if _, ok := this.IdUserMapping[followeeId]; !ok {
		this.IdUserMapping[followeeId] = &User{
			Id: followeeId,
		}
	}
	user := this.IdUserMapping[followerId]
	if _, ok := user.Followings[followeeId]; ok {
		return
	}
	if user.Followings == nil {
		user.Followings = make(map[int]*User)
	}
	user.Followings[followeeId] = this.IdUserMapping[followeeId]
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	user, ok := this.IdUserMapping[followerId]
	if !ok {
		return
	} else {
		if _, ok := user.Followings[followeeId]; ok {
			delete(user.Followings, followeeId)
		}
	}
}

type User struct {
	Id         int
	Followings map[int]*User
	TweetsHead *Tweet
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
