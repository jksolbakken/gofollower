package linkfollower

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

const maxRedirectDepth = 10

var metaRefreshPattern = regexp.MustCompile(`(?i)<(meta)\s+(http-equiv).*(content)=["']0;[ ]*(url)=(?P<Location>.*?)(["']*>)`)
var lnkdInPattern = regexp.MustCompile(`<a.*name="external_url_click".*>\s+(?P<Location>https?://.*\s+)</a>`)

type VisitResponse struct {
	VisitedURL     *url.URL
	IsRedirect     bool
	StatusCode     int
	Location       *url.URL
	AdditionalInfo string
}

type ResponseHandler func(resp VisitResponse)

type linkExtractor = func(html string) (*url.URL, error)

func Follow(startURL *url.URL, responseHandler ResponseHandler) error {
	u := prefixWithHttps(startURL)
	httpClient := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	for range maxRedirectDepth {
		response, err := visit(u, httpClient)
		if err != nil {
			return err
		}
		responseHandler(response)
		if !response.IsRedirect {
			return nil
		}
		u = response.Location
	}
	return fmt.Errorf("max redirect limit of %d exceeded", maxRedirectDepth)
}

func prefixWithHttps(u *url.URL) *url.URL {
	if u.Scheme == "" {
		u.Scheme = "https"
	}
	return u
}

func visit(site *url.URL, httpClient *http.Client) (VisitResponse, error) {
	resp, err := httpClient.Get(site.String())
	additional := ""
	if err != nil {
		return VisitResponse{}, err
	}
	redirectLocation, err := redirectByStatusCode(resp)
	if err != nil {
		return VisitResponse{}, err
	}
	if redirectLocation == nil && resp.StatusCode == 200 {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return VisitResponse{}, err
		}
		for _, extractor := range extractors {
			location, err := extractor(string(body))
			if err != nil {
				return VisitResponse{}, err
			}
			if location != nil {
				redirectLocation = location
				additional = "extracted from body"
				break
			}
		}

	}
	return VisitResponse{
		VisitedURL:     site,
		IsRedirect:     redirectLocation != nil,
		StatusCode:     resp.StatusCode,
		Location:       redirectLocation,
		AdditionalInfo: additional,
	}, nil
}

func redirectByStatusCode(resp *http.Response) (*url.URL, error) {
	isRedirect := resp.StatusCode == 301 ||
		resp.StatusCode == 302 ||
		resp.StatusCode == 303 ||
		resp.StatusCode == 307 ||
		resp.StatusCode == 308
	if isRedirect {
		return resp.Location()
	}
	return nil, nil
}

func redirectByMetaRefresh(input string) (*url.URL, error) {
	matches := metaRefreshPattern.FindStringSubmatch(input)
	if matches == nil {
		return nil, nil
	}
	locationIndex := metaRefreshPattern.SubexpIndex("Location")
	return url.Parse(matches[locationIndex])
}

func redirectByLnkdIn(input string) (*url.URL, error) {
	matches := lnkdInPattern.FindStringSubmatch(input)
	if matches == nil {
		return nil, nil
	}
	locationIndex := lnkdInPattern.SubexpIndex("Location")
	withoutTrailingSpaces := strings.TrimSpace(matches[locationIndex])
	return url.Parse(withoutTrailingSpaces)
}

var extractors = []linkExtractor{
	redirectByMetaRefresh,
	redirectByLnkdIn,
}
