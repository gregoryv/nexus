[![Build Status](https://travis-ci.org/gregoryv/nexus.svg?branch=master)](https://travis-ci.org/gregoryv/nexus)
[![codecov](https://codecov.io/gh/gregoryv/nexus/branch/master/graph/badge.svg)](https://codecov.io/gh/gregoryv/nexus)
[![Maintainability](https://api.codeclimate.com/v1/badges/df2736e1ac63580b49d7/maintainability)](https://codeclimate.com/github/gregoryv/nexus/maintainability)

[nexus](https://godoc.org/github.com/gregoryv/nexus) - package implements functions using nexus pattern

Primary use of [nexus pattern](https://www.7de.se/go-learn/SE/nexus_pattern.html) is in error handling.

## Quick start

    go get github.com/gregoryv/nexus

    func RenderSomething(w io.Writer) error {
	    var err *error
	    print := nexus.Print(w, err)
		print("My long")
		print("text")

		printf := nexus.Printf(err)
		printf("Hello, %s", "World")
		return *err
	}
