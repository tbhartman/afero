package afero

import afero "github.com/tbhartman/afero-lite/definition"

const BADFD = afero.BADFD
const FilePathSeparator = afero.FilePathSeparator

var ErrDestinationExists = afero.ErrDestinationExists
var ErrFileClosed = afero.ErrFileClosed
var ErrFileExists = afero.ErrFileExists
var ErrFileNotFound = afero.ErrFileNotFound
var ErrNoReadlink = afero.ErrNoReadlink
var ErrNoSymlink = afero.ErrNoSymlink
var ErrOutOfRange = afero.ErrOutOfRange
var ErrTooLarge = afero.ErrTooLarge

type DirsMerger = afero.DirsMerger
type File = afero.File
type Fs = afero.Fs

var NewOsFs = afero.NewOsFs

type LinkReader = afero.LinkReader
type Linker = afero.Linker
type Lstater = afero.Lstater
type OsFs = afero.OsFs
type Symlinker = afero.Symlinker
type UnionFile = afero.UnionFile
