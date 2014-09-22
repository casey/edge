graph
=====

A directed graph building service.

Imaginary API
-------------

* POST /A/B -> Increment the A -> B directed edge weight
* GET /A/B -> Get A -> B edge weight
* GET /A -> Get sum of outgoing edge weights from A
* GET /A/N -> Get Nth heavest edge from A
* GET /N/A -> Get Nth heavest edge to A
