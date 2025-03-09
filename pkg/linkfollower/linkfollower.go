package linkfollower

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
)

const maxRedirectDepth = 10

var metaRefreshPattern = regexp.MustCompile(`<(meta|META)\s+(http-equiv|HTTP-EQUIV).*(CONTENT|content)=["']0;[ ]*(URL|url)=(?P<Location>.*?)(["']*>)`)

type VisitResponse struct {
	IsRedirect bool
	StatusCode int
	Location   *url.URL
	Additional string
}

func Follow(startURL *url.URL, responseHandler func(visitedURL *url.URL, resp VisitResponse)) error {
	u := startURL
	httpClient := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	for i := 0; i < maxRedirectDepth; i++ {
		response, err := visit(u, httpClient)
		if err != nil {
			return err
		}
		responseHandler(u, response)
		if !response.IsRedirect {
			return nil
		}
		u = response.Location
	}
	return fmt.Errorf("max redirect limit of %d exceeded", maxRedirectDepth)
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
	if redirectLocation == nil && resp.Status == "200" {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return VisitResponse{}, err
		}
		redirectLocation, err = redirectByMetaRefresh(string(body))
		if err != nil {
			return VisitResponse{}, err
		}
		additional = "+ html meta refresh"
	}

	return VisitResponse{
		IsRedirect: redirectLocation != nil,
		StatusCode: resp.StatusCode,
		Location:   redirectLocation,
		Additional: additional,
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
