// Package fi - An interface exposing all exported methods of *os.File.
package fi

import (
	"io"
	"os"
	"syscall"
	"time"
)

// File is an interface representing the exported methods on *os.File
type File interface {
	// Comments on method names exist to link to the *os.File documentation

	// Readdir represents [os.File.Readdir]
	Readdir(n int) ([]os.FileInfo, error)
	// Readdirnames represents [os.File.Readdirnames]
	Readdirnames(n int) (names []string, err error)
	// ReadDir represents [os.File.ReadDir]
	ReadDir(n int) ([]os.DirEntry, error)
	// Close represents [os.File.Close]
	Close() error
	// Chown represents [os.File.Chown]
	Chown(uid int, gid int) error
	// Truncate represents [os.File.Truncate]
	Truncate(size int64) error
	// Sync represents [os.File.Sync]
	Sync() error
	// Chdir represents [os.File.Chdir]
	Chdir() error
	// Stat represents [os.File.Stat]
	Stat() (os.FileInfo, error)
	// Name represents [os.File.Name]
	Name() string
	// Read represents [os.File.Read]
	Read(b []byte) (n int, err error)
	// ReadAt represents [os.File.ReadAt]
	ReadAt(b []byte, off int64) (n int, err error)
	// ReadFrom represents [os.File.ReadFrom]
	ReadFrom(r io.Reader) (n int64, err error)
	// Write represents [os.File.Write]
	Write(b []byte) (n int, err error)
	// WriteAt represents [os.File.WriteAt]
	WriteAt(b []byte, off int64) (n int, err error)
	// Seek represents [os.File.Seek]
	Seek(offset int64, whence int) (ret int64, err error)
	// WriteString represents [os.File.WriteString]
	WriteString(s string) (n int, err error)
	// Chmod represents [os.File.Chmod]
	Chmod(mode os.FileMode) error
	// SetDeadline represents [os.File.SetDeadline]
	SetDeadline(t time.Time) error
	// SetReadDeadline represents [os.File.SetReadDeadline]
	SetReadDeadline(t time.Time) error
	// SetWriteDeadline represents [os.File.SetWriteDeadline]
	SetWriteDeadline(t time.Time) error
	// SyscallConn represents [os.File.SyscallConn]
	SyscallConn() (syscall.RawConn, error)
	// Fd represents [os.File.Fd]
	Fd() uintptr
}

var _ File = (*os.File)(nil)
