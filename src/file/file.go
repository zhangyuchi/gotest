// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package file

import (
	"os"
	"syscall"
)

type File struct {
	fd   int    // file descriptor number
	name string // file name at Open time
}

func newFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name}
}

var (
	Stdin  = newFile(syscall.Stdin, "/dev/stdin")
	Stdout = newFile(syscall.Stdout, "/dev/stdout")
	Stderr = newFile(syscall.Stderr, "/dev/stderr")
)

func Open(name string, mode int, perm uint32) (*File, error) {
	r, e := syscall.Open(name, mode, perm)
	return newFile(r, name), e
}

func (file *File) Close() error {
	if file == nil {
		return os.ErrInvalid
	}
	e := syscall.Close(file.fd)
	file.fd = -1 // so it can't be closed again
	return e
}

func (file *File) Read(b []byte) (int, error) {
	if file == nil {
		return -1, os.ErrInvalid
	}
	r, e := syscall.Read(file.fd, b)
	return int(r), e
}

func (file *File) ReadLine(b []byte) (int, error) {
	if file == nil {
		return -1, os.ErrInvalid
	}
	r, e := syscall.Read(file.fd, b)
	return int(r), e
}

func (file *File) Write(b []byte) (int, error) {
	if file == nil {
		return -1, os.ErrInvalid
	}
	r, e := syscall.Write(file.fd, b)
	return int(r), e
}

func (file *File) String() string {
	return file.name
}
