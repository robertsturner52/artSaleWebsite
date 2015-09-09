package heatherandersonwilderness

import (
    "net/http"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static/"+r.URL.Path)
}