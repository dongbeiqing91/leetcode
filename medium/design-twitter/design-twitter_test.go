package designtwitter

import (
	"reflect"
	"testing"
)

func TestExample(t *testing.T) {
	tw := Constructor()
	tw.PostTweet(1, 5)
	if got := tw.GetNewsFeed(1); !reflect.DeepEqual(got, []int{5}) {
		t.Fatalf("feed after post = %v, want [5]", got)
	}
	tw.Follow(1, 2)
	tw.PostTweet(2, 6)
	if got := tw.GetNewsFeed(1); !reflect.DeepEqual(got, []int{6, 5}) {
		t.Fatalf("feed after follow = %v, want [6 5]", got)
	}
	tw.Unfollow(1, 2)
	if got := tw.GetNewsFeed(1); !reflect.DeepEqual(got, []int{5}) {
		t.Fatalf("feed after unfollow = %v, want [5]", got)
	}
}

func TestUnfollowNonFollowed(t *testing.T) {
	tw := Constructor()
	tw.PostTweet(1, 5)
	tw.Follow(1, 2)
	tw.PostTweet(2, 6)
	// unfollow someone never followed should be a no-op
	tw.Unfollow(1, 3)
	if got := tw.GetNewsFeed(1); !reflect.DeepEqual(got, []int{6, 5}) {
		t.Fatalf("feed = %v, want [6 5]", got)
	}
}

func TestFollowSelf(t *testing.T) {
	tw := Constructor()
	tw.Follow(1, 1)
	tw.PostTweet(1, 5)
	if got := tw.GetNewsFeed(1); !reflect.DeepEqual(got, []int{5}) {
		t.Fatalf("feed = %v, want [5]", got)
	}
}

func TestTenLimit(t *testing.T) {
	tw := Constructor()
	for i := 1; i <= 12; i++ {
		tw.PostTweet(1, i)
	}
	want := []int{12, 11, 10, 9, 8, 7, 6, 5, 4, 3}
	if got := tw.GetNewsFeed(1); !reflect.DeepEqual(got, want) {
		t.Fatalf("feed = %v, want %v", got, want)
	}
}
