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
  * `util` package created
    - moves most utility functions here
	- file `util/afero.go` contains limited import of main afero interfaces
	- file `unionFile.go` split and `copy*` functions moved to `util`
  * `mem` package now contains `NewMemMapFs`
  * `httpfs` package created, with `NewHttpFs` moved here

Thus, the root package only defines the interfaces and allows using
the OsFs via `NewOsFs`.  Use `httpfs`, knowing that `net/http` will be imported
and increase binary size unless you need it.  Use `util` for all
other afero utilities.

Versioning
----------

Versioning will follow major and minor releases of [afero]("github.com/spf13/afero"),
*but the patch version will update independently as patches to this repo are required!*

Update instructions
-------------------

This is all manual right now :(

 * remove `_test.go` and `mock.go` files
 * create `httpfs` and `util` directories
 * split `FilePathSeparator` from `util.go` and put remaining in `util`
 * split `copy*` functions from `unionFile.go` and put in `util`
 * move `httpfs.go` to `httpfs`
 * move `memmap.go` to `mem`
 * move all other root go files to `util`
 * rename packages accordingly (from `package afero`)
 * add `afero.go` to `util` and define interfaces from root
   - e.g. `var Fs = afero.Fs`
 * fix build errors (e.g. `httpfs.go` and `memmap.go` will need `afero.` prefixes for interfaces)
