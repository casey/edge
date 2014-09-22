package app

import "net/http"
import "regexp"
import "fmt"
import "appengine"

import . "flotilla"

var path_re = regexp.MustCompile(`^/(?P<a>[a-zA-Z0-9._-]{1,64})/(?P<b>[a-zA-Z0-9._-]{1,64})$`)

func init() {
  Debug(true)
  Handle("/").Get(get).Post(post).Options(options)
}

func options(r *http.Request) {
  Status(http.StatusOK)
}

func get(r *http.Request) {
  c := appengine.NewContext(r)
  match := Components(path_re, r.URL.Path)
  Ensure(match != nil, http.StatusForbidden)
  w, e := weight(c, match["a"], match["b"])
  Check(e)
  Body(http.StatusOK, fmt.Sprintf("%v", w), "text/plain")
}

func post(r *http.Request) {
  c := appengine.NewContext(r)
  match := Components(path_re, r.URL.Path)
  Ensure(match != nil, http.StatusForbidden)
  w, e := increment(c, match["a"], match["b"])
  Check(e)
  Body(http.StatusOK, fmt.Sprintf("%v", w), "text/plain")
}
