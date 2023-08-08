afero-lite
==========

A *lightweight* FileSystem Abstraction System for Go

Modifications
-------------

This is [afero](http://github.com/spf13/afero), reorganized to avoid bloated
binaries when a subset of features is used (aka when `net/http` is not needed).
The following modifications are made:

  * `"github.com/spf13/afero` replaced with `"github.com/tbhartman/afero/lite`
  * all root go files are moved to a `lite` package, except for `httpFs.go`
  * new aliases are created at the root level to reference exported symbols
    in `lite`

Usage
-----

Unless you need `NewHttpFs`, use the `lite` package, which is imported as
`afero`.

```go
package main

import "github.com/tbhartman/afero/lite"

func main() {
	var myfs afero.Fs
	myfs = afero.NewOsFs()
	myfs.Stat("myfile")
}
```

If you need `httpFs` functionality, import from root:

```go
package main

import (
	"github.com/tbhartman/afero"
)

func main() {
	var myfs afero.Fs
	myfs = afero.NewHttpFs(afero.NewMemMapFs())
}
```

Versioning
----------

Versioning will follow major and minor releases of [afero]("github.com/spf13/afero"),
*but the patch version will update independently as patches to this repo are required!*
