package app

import "net/http"
import "regexp"
import "appengine"

import . "flotilla"

var path_re = regexp.MustCompile(`^/(?P<a>[a-zA-Z0-9._-]{1,64})(?P<b>[a-zA-Z0-9._-]{1,64})$`)

func init() {
  Handle("/").Get(get).Post(post).Options(options)
}

func options(r *http.Request) {
  Status(http.StatusOK)
}

func get(r *http.Request) {
  _ = appengine.NewContext(r)
  match := Components(path_re, r.URL.Path)
  Ensure(match != nil, http.StatusForbidden)
  Status(http.StatusNotImplemented)
}

func post(r *http.Request) {
  _ = appengine.NewContext(r)
  match := Components(path_re, r.URL.Path)
  Ensure(match != nil, http.StatusForbidden)
  Status(http.StatusNotImplemented)
}
