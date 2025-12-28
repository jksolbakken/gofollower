package linkfollower

import (
	"net/url"
	"testing"
)

func TestMetaRefreshTagYieldsURL(t *testing.T) {
	expected, _ := url.Parse("http://yolo.test")
	actual, _ := redirectByMetaRefresh(`<META http-equiv="refresh" content="0; url=http://yolo.test">`)
	if actual.String() != expected.String() {
		t.Errorf("URL was incorrect, got: %v, want: %v.", actual, expected)
	}
}

func TestNoMetaRefreshTagYieldsNil(t *testing.T) {
	input := "bla bla unrelated"
	actual, _ := redirectByMetaRefresh(`bogus input`)
	if actual != nil {
		t.Errorf("%s does not contain a valid URL", input)
	}
}

func TestMetaRefreshIsCaseInsensitive(t *testing.T) {
	expected, _ := url.Parse("http://yolo.test")
	actual, _ := redirectByMetaRefresh(`<meta HTTP-EQUIV="refresh" CONTENT="0; url=http://yolo.test">`)
	if actual.String() != expected.String() {
		t.Errorf("URL was incorrect, got: %v, want: %v.", actual, expected)
	}
}

func TestLinkedInYieldsURL(t *testing.T) {
	input := `<a class="artdeco-button artdeco-button--tertiary" data-tracking-control-name="external_url_click" data-tracking-will-navigate href="https://85340.webcruiter.no/Main2/Recruit/Public/4895658382?language=nb&amp;link_source_id=0">
                https://85340.webcruiter.no/Main2/Recruit/Public/4895658382?language=nb&amp;link_source_id=0
            </a>`

	expected, _ := url.Parse("https://85340.webcruiter.no/Main2/Recruit/Public/4895658382?language=nb&amp;link_source_id=0")
	actual, _ := redirectByLinkedIn(input)
	if actual.String() != expected.String() {
		t.Errorf("URL was incorrect, got: %v, want: %v.", actual, expected)
	}
}

func TestNoLinkedInAnchorYieldsNil(t *testing.T) {
	input := "bla bla unrelated"
	actual, _ := redirectByLinkedIn(input)
	if actual != nil {
		t.Errorf("%s does not contain a valid URL", input)
	}
}
