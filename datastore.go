package app

import "appengine"
import "appengine/datastore"

type Edge struct {
  Weight int64 `datastore:"weight",noindex`
}

func datastoreKey(c appengine.Context, a string, b string) *datastore.Key {
  return datastore.NewKey(c, "Edge", a + "/" + b, 0, nil)
}

func weight(c appengine.Context, a string, b string) (int64, error) {
  key := datastoreKey(c, a, b)
  edge := Edge{}
  e := datastore.Get(c, key, &edge)
  if e == datastore.ErrNoSuchEntity {
    return 0, nil
  } else if e != nil {
    return 0, nil
  } else {
    return edge.Weight, nil
  }
}

func increment(c appengine.Context, a string, b string) (int64, error) {
  var weight int64
  e := datastore.RunInTransaction(c, func(c appengine.Context) error {
    key := datastoreKey(c, a, b)
    edge := Edge{}
    e := datastore.Get(c, key, &edge)
    if e == datastore.ErrNoSuchEntity {
      edge.Weight = 0
    } else if e != nil {
      return e
    }
    edge.Weight++
    weight = edge.Weight
    _, e = datastore.Put(c, key, &edge)
    return e
  }, nil)

  if e != nil {
    return 0, e
  }

  return weight, nil
}
