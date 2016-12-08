package medium

import "testing"

func TestMediumCrawler_Parse(t *testing.T) {
	crawler := NewCrawler("https://medium.com/browse/top")
	feed,err := crawler.Parse()
	if err != nil {
		t.Error(err)
	}
	if len(feed.Items) <= 0 {
		t.Error("Can not get items")
	}
}
