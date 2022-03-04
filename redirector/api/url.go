package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"route"
)

var routes map[string]string

func init() {
	_ = json.Unmarshal(route.Data, &routes)
}

func UrlHandler(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Path) <= 1 {
		_, _ = fmt.Fprintf(w, "Invaid short name")
		return
	}
	if url := routes[r.URL.Path[1:]]; url != "" {
		http.Redirect(w, r, url, http.StatusMovedPermanently)
	} else {
		_, _ = fmt.Fprintf(w, "Invaid short name")
	}
	return
}
