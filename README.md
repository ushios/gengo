ushios/gengo
=============
[![Build Status](https://travis-ci.org/ushios/gengo.svg?branch=master)](https://travis-ci.org/ushios/gengo) [![Coverage Status](https://coveralls.io/repos/github/ushios/gengo/badge.svg?branch=master)](https://coveralls.io/github/ushios/gengo?branch=master)

Get gengo

Instalation
-------------

```shell
$ go get github.com/ushios/gengo
```

Examples
--------

```go
g := gengo.Now()

fmt.Println(g)
// Output:
// 平成
```

```go
t, _ := time.Parse("2006-01-02 15:04:05 -0700", "1912-07-29 23:59:59 +0900")
g := gengo.At(t)

fmt.Println(g)
// Output:
// 明治
```
