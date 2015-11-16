# bom
## byte order mark

[![GoDoc](https://godoc.org/github.com/spkg/bom?status.svg)](https://godoc.org/github.com/spkg/bom)
[![Build Status](https://travis-ci.org/spkg/bom.svg?branch=master)](https://travis-ci.org/spkg/bom)
[![license](http://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://raw.githubusercontent.com/spkg/bom/master/LICENSE.md)
[![Coverage](http://gocover.io/_badge/github.com/spkg/bom)](http://gocover.io/github.com/spkg/bom)

The `bom` package provides a convenient way to strip UTF-8 byte order marks (BOM) from the
beginning of a byte slice or an `io.Reader`.

The Go standard library provides no support for UTF-8 byte order marks, and it looks like it never will. 
To quote Andy Balholm in the discussion on this issue at https://groups.google.com/forum/#!topic/golang-nuts/OToNIPdfkks

>  The Go team includes the original designers of UTF-8, and they consider BOMs an aBOMination.
  They are reluctant to do anything to make life easier for people who use BOMs. :-)

>  (Although they did make the compiler accept source files with BOMs, if I remember right.)

In the same discussion thread another participant makes the comment that it should not be difficult to write 
an `io.Reader` that eats the BOM.

It isn't difficult, and here it is.

