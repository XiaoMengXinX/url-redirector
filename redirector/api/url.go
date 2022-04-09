package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"route"
)

type routeData struct {
	Scope string            `json:"scope"`
	Rules map[string]string `json:"rules"`
}

var routes []routeData

func init() {
	_ = json.Unmarshal(route.Data, &routes)
}

func UrlHandler(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Path) <= 1 {
		_, _ = fmt.Fprintf(w, "Invaid short name")
		return
	}
	for _, rt := range routes {
		scope := regexp.MustCompile(rt.Scope)
		if scope.MatchString(r.URL.Host) {
			if url := rt.Rules[r.URL.Path[1:]]; url != "" {
				http.Redirect(w, r, url, http.StatusMovedPermanently)
				return
			} else {
				_, _ = fmt.Fprintf(w, "Invaid short name")
				return
			}
		}
	}
}
