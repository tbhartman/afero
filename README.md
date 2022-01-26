afero-lite
==========

A *lightweight* FileSystem Abstraction System for Go

Modifications
-------------

This is [afero]("github.com/spf13/afero"), reorganized to avoid
bloated binaries when a subset of features is used.  The following
modifications are made:

  * tests removed
    - allfiles ending `_test.go`, as well as `gcs_mocks.go`
	- count on afero to do the testing
  * `full` package created
    - all of afero moved here except that related to `MemMapFs` and `HttpFs`
  * `mem` package now contains `NewMemMapFs`
  * `httpfs` package created, with `NewHttpFs` moved here

Thus, the root package only defines the interfaces and allows using
the OsFs via `NewOsFs`.  This should keep the amount of imported code to
a bare minimum if all you need is to abstract `os` usage.
Use `httpfs`, knowing that `net/http` will be imported
and increase binary size unless you need it.  Use `full` for all
other afero utilities.

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

If you need other functionality, import from `full`, and use
as `afero`:

```go
package main

import (
	"github.com/tbhartman/afero-lite/full"
)

func main() {
	var myfs afero.Fs
	myfs = afero.NewOsFs()
	myfs = afero.NewBasePath(myfs, "basepath")
	myfs.Stat("myfile")
}
```

If you need a `MemMapFs`, import from `mem`:

```go
package main

import (
	"github.com/tbhartman/afero-lite/mem"
)

func main() {
	_ = mem.NewMemMapFs()
}
```

If you need a `HttpFs`, import from `httpfs`:

```go
package main

import (
	"github.com/tbhartman/afero-lite/httpfs"
)

func main() {
	_ = httpfs.NewHttpFs()
}
```

Versioning
----------

Versioning will follow major and minor releases of [afero]("github.com/spf13/afero"),
*but the patch version will update independently as patches to this repo are required!*

Update instructions
-------------------

This is all manual right now :(

 * remove `_test.go` and `mock.go` files
 * create `httpfs` and `full` directories
 * move `httpfs.go` to `httpfs`
 * move `memmap.go` to `mem`
 * move all other root go files to `full`
 * copy back some files to root:
   - `afero.go`
   - `const_bsds.go`
   - `const_win_unix.go`
   - `lstater.go`
   - `os.go`
   - `symlink.go`
   - `unionFile.go`
   - `util.go`
 * keep only `FilePathSeparator` in root `util.go`
 * remove `copy*` functions from root `unionFile.go`
 * fix build errors
   - `httpfs.go` will need to import and use `afero`
   - `memmap.go` will need to import and use `afero`
