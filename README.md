ushios/gengo
=============
[![Build Status](https://travis-ci.org/ushios/gengo.svg?branch=master)](https://travis-ci.org/ushios/gengo) [![Coverage Status](https://coveralls.io/repos/github/ushios/gengo/badge.svg?branch=master)](https://coveralls.io/github/ushios/gengo?branch=master)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fushios%2Fgengo.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fushios%2Fgengo?ref=badge_shield)

Get gengo

Instalation
-------------

```bash
$ go get github.com/ushios/gengo
```

Documentation
-------------

[![GoDoc](https://godoc.org/github.com/ushios/gengo?status.svg)](https://godoc.org/github.com/ushios/gengo)

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


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fushios%2Fgengo.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fushios%2Fgengo?ref=badge_large)