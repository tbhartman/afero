afero-lite
==========

A *lightweight* FileSystem Abstraction System for Go

Modifications
-------------

This is [afero](http://github.com/spf13/afero), reorganized to avoid bloated
binaries when a subset of features is used (aka when `net/http` is not needed).
The following modifications are made:

  * `"github.com/spf13/afero` replaced with `"github.com/tbhartman/afero-lite`
  * `httpfs` package created, with `httpFs.go` moved here
    - need to import from base package for `Fs` and `File` definitions

Usage
-----

The top-level package is simply named `afero`:

```go
package main

import "github.com/tbhartman/afero-lite"

func main() {
	var myfs afero.Fs
	myfs = afero.NewOsFs()
	myfs.Stat("myfile")
}
```

If you need `httpFs` functionality, import from `httpfs`:

```go
package main

import (
	"github.com/tbhartman/afero-lite"
	"github.com/tbhartman/afero-lite/httpfs"
)

func main() {
	var myfs afero.Fs
	myfs = httpfs.NewHttpFs(afero.NewMemMapFs())
}
```

Versioning
----------

Versioning will follow major and minor releases of [afero]("github.com/spf13/afero"),
*but the patch version will update independently as patches to this repo are required!*
