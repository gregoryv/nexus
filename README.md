[![Build Status](https://travis-ci.org/gregoryv/nexus.svg?branch=master)](https://travis-ci.org/gregoryv/nexus)
[![codecov](https://codecov.io/gh/gregoryv/nexus/branch/master/graph/badge.svg)](https://codecov.io/gh/gregoryv/nexus)
[![Maintainability](https://api.codeclimate.com/v1/badges/df2736e1ac63580b49d7/maintainability)](https://codeclimate.com/github/gregoryv/nexus/maintainability)

[nexus](https://godoc.org/github.com/gregoryv/nexus) - package implements functions using nexus pattern

Primary use of [nexus pattern](https://www.7de.se/go-learn/SE/nexus_pattern.html) is in error handling.

## Quick start

    go get github.com/gregoryv/nexus

    func RenderSomething(w io.Writer) error {
        p, err := nexus.NewPrinter(os.Stdout)
        p.Print("My long")
        p.Println("text")
        p.Printf("Hello, %s", "World")
        return *err
    }
