package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"route"
)

var routes map[string]string

type resData struct {
	Token string `json:"token"`
	Error string `json:"error"`
}

func init() {
	_ = json.Unmarshal(route.Data, &routes)
}

func UrlHandler(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Path) <= 1 {
		_, _ = fmt.Fprintf(w, responseJson(resData{Error: "Invaid short name"}))
		return
	}
	if url := routes[r.URL.Path[1:]]; url != "" {
		http.Redirect(w, r, url, http.StatusMovedPermanently)
	} else {
		_, _ = fmt.Fprintf(w, responseJson(resData{Error: "Invaid short name"}))
	}
	return
}

func responseJson(r resData) string {
	b, _ := json.Marshal(r)
	return string(b)
}
