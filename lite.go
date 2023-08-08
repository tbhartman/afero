// Code generated with liteGen.go DO NOT EDIT.
package afero

import (
	"errors"

	lite "github.com/tbhartman/afero/lite"
)

// manually alias ErrOutOfRange
var ErrOutOfRange = errors.New("out of range")

// aliasing types from lite package
type Afero = lite.Afero
type BasePathFile = lite.BasePathFile
type BasePathFs = lite.BasePathFs
type CacheOnReadFs = lite.CacheOnReadFs
type CopyOnWriteFs = lite.CopyOnWriteFs
type DirsMerger = lite.DirsMerger
type File = lite.File
type FromIOFS = lite.FromIOFS
type Fs = lite.Fs
type IOFS = lite.IOFS
type LinkReader = lite.LinkReader
type Linker = lite.Linker
type Lstater = lite.Lstater
type MemMapFs = lite.MemMapFs
type OsFs = lite.OsFs
type ReadOnlyFs = lite.ReadOnlyFs
type RegexpFile = lite.RegexpFile
type RegexpFs = lite.RegexpFs
type Symlinker = lite.Symlinker
type UnionFile = lite.UnionFile

// aliasing constants from lite package
const BADFD = lite.BADFD
const FilePathSeparator = lite.FilePathSeparator

// aliasing variables from lite package
var ErrFileClosed = lite.ErrFileClosed
var ErrNoReadlink = lite.ErrNoReadlink
var ErrNoSymlink = lite.ErrNoSymlink

// aliasing functions from lite package
var DirExists = lite.DirExists
var Exists = lite.Exists
var FileContainsAnyBytes = lite.FileContainsAnyBytes
var FileContainsBytes = lite.FileContainsBytes
var FullBaseFsPath = lite.FullBaseFsPath
var GetTempDir = lite.GetTempDir
var Glob = lite.Glob
var IsDir = lite.IsDir
var IsEmpty = lite.IsEmpty
var NeuterAccents = lite.NeuterAccents
var NewBasePathFs = lite.NewBasePathFs
var NewCacheOnReadFs = lite.NewCacheOnReadFs
var NewCopyOnWriteFs = lite.NewCopyOnWriteFs
var NewIOFS = lite.NewIOFS
var NewMemMapFs = lite.NewMemMapFs
var NewOsFs = lite.NewOsFs
var NewReadOnlyFs = lite.NewReadOnlyFs
var NewRegexpFs = lite.NewRegexpFs
var ReadAll = lite.ReadAll
var ReadDir = lite.ReadDir
var ReadFile = lite.ReadFile
var SafeWriteReader = lite.SafeWriteReader
var TempDir = lite.TempDir
var TempFile = lite.TempFile
var UnicodeSanitize = lite.UnicodeSanitize
var Walk = lite.Walk
var WriteFile = lite.WriteFile
var WriteReader = lite.WriteReader
