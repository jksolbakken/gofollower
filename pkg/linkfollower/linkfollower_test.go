package linkfollower

import (
	"net/url"
	"testing"
)

func TestExtractingFromMetaRefresh(t *testing.T) {
	expected, _ := url.Parse("http://yolo.test")
	actual, _ := redirectByMetaRefresh(`<META http-equiv="refresh" content="0; url=http://yolo.test">`)
	if actual.String() != expected.String() {
		t.Errorf("URL was incorrect, got: %v, want: %v.", actual, expected)
	}
}

func TestExtractingFromMetaRefresh2(t *testing.T) {
	input := "bla bla unrelated"
	actual, _ := redirectByMetaRefresh(`bogus input`)
	if actual != nil {
		t.Errorf("%s does not containn a valid URL", input)
	}
}
