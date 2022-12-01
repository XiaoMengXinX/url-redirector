package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"api/config"
)

type routeData struct {
	Scope string            `json:"scope"`
	Rules map[string]string `json:"rules"`
}

var routes []routeData

func init() {
	_ = json.Unmarshal(config.Data, &routes)
}

func UrlHandler(w http.ResponseWriter, r *http.Request) {
	for _, rt := range routes {
		scope := regexp.MustCompile(rt.Scope)
		if scope.MatchString(r.URL.Host) {
			if rt.Rules["/"] != "" && r.URL.Path == "" {
				http.Redirect(w, r, rt.Rules["/"], http.StatusFound)
				return
			}
			if url := rt.Rules[r.URL.Path[1:]]; url != "" {
				http.Redirect(w, r, url, http.StatusFound)
				return
			} else {
				_, _ = fmt.Fprintf(w, "Invaid short name")
				return
			}
		}
	}
}
