package linkfollower

import (
	"fmt"
	"net/http"
	"net/url"
)

const maxRedirectDepth = 10

type VisitResponse struct {
	IsRedirect bool
	StatusCode int
	Location   *url.URL
}

func Follow(startURL *url.URL, responseHandler func(visitedURL *url.URL, resp VisitResponse)) error {
	u := startURL
	for i := 0; i < maxRedirectDepth; i++ {
		response, err := visit(u)
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

func visit(site *url.URL) (VisitResponse, error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Get(site.String())
	if err != nil {
		return VisitResponse{}, err
	}
	isRedirect := isRedirect(resp)
	var location *url.URL
	if isRedirect {
		location, err = resp.Location()
		if err != nil {
			return VisitResponse{}, err
		}
	}
	return VisitResponse{
		IsRedirect: isRedirect,
		StatusCode: resp.StatusCode,
		Location:   location,
	}, nil
}

func isRedirect(resp *http.Response) bool {
	return resp.StatusCode == 301 ||
		resp.StatusCode == 302 ||
		resp.StatusCode == 303 ||
		resp.StatusCode == 307 ||
		resp.StatusCode == 308
}
