package linkfollower

import (
	"net/url"
	"testing"
)

func TestStuff(t *testing.T) {
	url, _ := url.Parse("https://www.example.com?a=b&utm_something=yolo")
	StripTracking(url)
	expected := "https://www.example.com?a=b"
	actual := url.String()
	if expected != actual {
		t.Errorf("want %s, got %s", expected, actual)
	}
}

