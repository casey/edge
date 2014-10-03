edge
====

Directed graph building service

[Test instance here.](http://rodarmor-edge.appspot.com)


API
---

A vertex name matches `/[a-zA-Z0-9._-]{1,64}/`.

* POST /A/B -> Increment the weight of the edge from vertex A to vertex B.
* GET /A/B -> Get weight of the edge from vertex A to vertex B. 0 if edge has never been incremented.

```
> curl -X POST 'http://rodarmor-edge.appspot.com/hello/world' --data ''
1
> curl -X POST 'http://rodarmor-edge.appspot.com/hello/world' --data ''
2
> curl -X POST 'http://rodarmor-edge.appspot.com/hello/world' --data ''
3
> curl -X GET 'http://rodarmor-edge.appspot.com/hello/world'
3
```
