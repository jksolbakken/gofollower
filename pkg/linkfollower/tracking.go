package linkfollower

import (
	"net/url"
	"strings"
)

var trackingParams = []string{
	"utm_",
	"mc_",
	"oly_",
	"_hs",
	"fbclid",
	"igshid",
	"mkt_tok",
	"otc",
	"wickedid",
	"_openstat",
	"yclid",
	"ICID",
	"rb_clickid",
}

func StripTracking(u *url.URL) {
	q := u.Query()
	var trackingParams []string
	for paramName := range q {
		if isTracking(paramName) {
			trackingParams = append(trackingParams, paramName)
		}
	}
	for _, tp := range trackingParams {
		delete(q, tp)
	}
	u.RawQuery = q.Encode()
}

func isTracking(queryParam string) bool {
	for _, p := range trackingParams {
		if strings.HasPrefix(queryParam, p) {
			return true
		}
	}
	return false
}
