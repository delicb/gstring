[![Build Status](https://travis-ci.org/delicb/gstring.svg?branch=master)](https://travis-ci.org/delicb/gstring)
[![Coverage](https://img.shields.io/badge/coverage-87.0%-brightgreen.svg)](http://gocover.io/github.com/delicb/gstring)
[![GoDoc](http://godoc.org/github.com/delicb/gstring?status.png)](http://godoc.org/github.com/delicb/gstring)
# gstring
Python like string formatting for golang.

## Why?
Golang already has great features for string formatting (entire [fmt](https://golang.org/pkg/fmt/) package) I missed named parameters for string formatting. Intention of this library is to mimic subset of features that python's format function offers.

## Example
Simple way of formatting string is using something like this:

```go
// outputs "some value"
gstring.Printm("{key}", map[string]interface{}{"key": "some value"})
```

If bracked is needed in string it can be created by escaping (two brackets). Also, single key can appear multiple times in format string:

```go
// outpits "Bracket {, }, key value, key value, key value"
gstring.Printm("Bracket {{, }}, {key}, {key}, {key}", map[string]interaface{}{"key", "key value"})
```

All standard formatting options from [fmt](https://golang.org/pkg/fmt) work. To specify them, add colon after key name and specify fmt package compatible formatting optinos. For example:

```go
// outpus "3.14"
gstring.Printm("{pi:%.2f}", map[string]interface{}{"pi": math.Pi})
```

## Installation
Run `go get github.com/delicb/gstring` from command line.
