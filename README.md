# bom
## strip UTF-8 byte order marks

[![GoDoc](https://pkg.go.dev/badge/github.com/spkg/bom)](https://godoc.org/github.com/spkg/bom)
[![License](http://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://raw.githubusercontent.com/spkg/bom/master/LICENSE.md)
[![GoReportCard](https://goreportcard.com/badge/github.com/spkg/bom)](http://goreportcard.com/report/spkg/bom)


The `bom` package provides a convenient way to strip [UTF-8 byte order marks](https://en.wikipedia.org/wiki/Byte_order_mark#UTF-8)
(BOM) from the beginning of a byte slice or an `io.Reader`.

The Unicode Standard defines UTF-8 byte order marks as the byte sequence `0xEF,0xBB,0xBF`, but neither requires nor recommends their use.
The Go standard library provides no support for UTF-8 byte order marks, and it looks like it never will. To quote Andy Balholm in the
discussion on this issue at https://groups.google.com/forum/#!topic/golang-nuts/OToNIPdfkks

>  The Go team includes the original designers of UTF-8, and they consider BOMs an aBOMination.
  They are reluctant to do anything to make life easier for people who use BOMs. :-)

>  (Although they did make the compiler accept source files with BOMs, if I remember right.)

In the same discussion thread another participant makes the comment that it should not be difficult to write
an `io.Reader` that eats the BOM.

It isn't difficult, and here is one simple implementation.

