# go-iid

Package **iid** provides quazi‐ monotonically‐increasing unique‐identifiers.

The serialized form of the **IID** is safe to use as a _file_ or _directory_ name.


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-iid

[![GoDoc](https://godoc.org/github.com/reiver/go-iid?status.svg)](https://godoc.org/github.com/reiver/go-iid)


## Example

Here is an example of using `package iid`:
```go
var id iid.IID = iid.Generate()
```
