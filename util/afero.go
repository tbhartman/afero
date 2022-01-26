package util

import "github.com/tbhartman/afero-lite"

type Fs = afero.Fs
type File = afero.File
type Linker = afero.Linker
type Lstater = afero.Lstater
type LinkReader = afero.LinkReader
type UnionFile = afero.UnionFile
type Afero afero.Afero

var ErrNoSymlink = afero.ErrNoSymlink
var ErrNoReadlink = afero.ErrNoReadlink
var ErrFileExists = afero.ErrFileExists
var ErrFileNotFound = afero.ErrFileNotFound
var FilePathSeparator = afero.FilePathSeparator
