package util

import (
	"io"
	"os"
	"path/filepath"
	"syscall"
)

func copyFile(base Fs, layer Fs, name string, bfh File) error {
	// First make sure the directory exists
	exists, err := Exists(layer, filepath.Dir(name))
	if err != nil {
		return err
	}
	if !exists {
		err = layer.MkdirAll(filepath.Dir(name), 0777) // FIXME?
		if err != nil {
			return err
		}
	}

	// Create the file on the overlay
	lfh, err := layer.Create(name)
	if err != nil {
		return err
	}
	n, err := io.Copy(lfh, bfh)
	if err != nil {
		// If anything fails, clean up the file
		layer.Remove(name)
		lfh.Close()
		return err
	}

	bfi, err := bfh.Stat()
	if err != nil || bfi.Size() != n {
		layer.Remove(name)
		lfh.Close()
		return syscall.EIO
	}

	err = lfh.Close()
	if err != nil {
		layer.Remove(name)
		lfh.Close()
		return err
	}
	return layer.Chtimes(name, bfi.ModTime(), bfi.ModTime())
}

func copyToLayer(base Fs, layer Fs, name string) error {
	bfh, err := base.Open(name)
	if err != nil {
		return err
	}
	defer bfh.Close()

	return copyFile(base, layer, name, bfh)
}

func copyFileToLayer(base Fs, layer Fs, name string, flag int, perm os.FileMode) error {
	bfh, err := base.OpenFile(name, flag, perm)
	if err != nil {
		return err
	}
	defer bfh.Close()

	return copyFile(base, layer, name, bfh)
}
