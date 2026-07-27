package linkfollower

import (
	"fmt"
	"net/url"
	"testing"
)

func TestBadParamsAreBlocked(t *testing.T) {
	for _, badParam := range trackingParams {
		u, _ := url.Parse(fmt.Sprintf("https://www.example.com?a=b&%s=yolo", badParam))
		StripTracking(u)
		expected := "https://www.example.com?a=b"
		actual := u.String()
		if expected != actual {
			t.Errorf("want %s, got %s", expected, actual)
		}
	}
}

func TestBenignParamsAreKeptAsIs(t *testing.T) {
	benignUrl := "https://www.example.com?a=b&c=d"
	u, _ := url.Parse(benignUrl)
	StripTracking(u)
	expected := benignUrl
	actual := u.String()
	if expected != actual {
		t.Errorf("want %s, got %s", expected, actual)
	}
}
