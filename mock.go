package fi

import (
	"errors"
	"io"
	"os"
	"syscall"
	"time"
)

// MockFileFromOS returns a *MockFile that delegates every method to the provided *os.File.
// Methods can then be overridden to customise behaviour, e.g. causing Write() to fail.
func MockFileFromOS(file *os.File) *MockFile {
	return &MockFile{
		ReaddirFn:          file.Readdir,
		ReaddirnamesFn:     file.Readdirnames,
		ReadDirFn:          file.ReadDir,
		CloseFn:            file.Close,
		ChownFn:            file.Chown,
		TruncateFn:         file.Truncate,
		SyncFn:             file.Sync,
		ChdirFn:            file.Chdir,
		StatFn:             file.Stat,
		NameFn:             file.Name,
		ReadFn:             file.Read,
		ReadAtFn:           file.ReadAt,
		ReadFromFn:         file.ReadFrom,
		WriteFn:            file.Write,
		WriteAtFn:          file.WriteAt,
		SeekFn:             file.Seek,
		WriteStringFn:      file.WriteString,
		ChmodFn:            file.Chmod,
		SetDeadlineFn:      file.SetDeadline,
		SetReadDeadlineFn:  file.SetReadDeadline,
		SetWriteDeadlineFn: file.SetWriteDeadline,
		SyscallConnFn:      file.SyscallConn,
		FdFn:               file.Fd,
	}
}

type MockFile struct {
	ReaddirFn          func(n int) ([]os.FileInfo, error)
	ReaddirnamesFn     func(n int) (names []string, err error)
	ReadDirFn          func(n int) ([]os.DirEntry, error)
	CloseFn            func() error
	ChownFn            func(uid int, gid int) error
	TruncateFn         func(size int64) error
	SyncFn             func() error
	ChdirFn            func() error
	StatFn             func() (os.FileInfo, error)
	NameFn             func() string
	ReadFn             func(b []byte) (n int, err error)
	ReadAtFn           func(b []byte, off int64) (n int, err error)
	ReadFromFn         func(r io.Reader) (n int64, err error)
	WriteFn            func(b []byte) (n int, err error)
	WriteAtFn          func(b []byte, off int64) (n int, err error)
	SeekFn             func(offset int64, whence int) (ret int64, err error)
	WriteStringFn      func(s string) (n int, err error)
	ChmodFn            func(mode os.FileMode) error
	SetDeadlineFn      func(t time.Time) error
	SetReadDeadlineFn  func(t time.Time) error
	SetWriteDeadlineFn func(t time.Time) error
	SyscallConnFn      func() (syscall.RawConn, error)
	FdFn               func() uintptr
}

var _ File = (*MockFile)(nil)

var ErrMockNotImplemented = errors.New("not implemented by mock")

func (m *MockFile) Readdir(n int) ([]os.FileInfo, error) {
	if m.ReaddirFn == nil {
		return nil, ErrMockNotImplemented
	}
	return m.ReaddirFn(n)
}

func (m *MockFile) Readdirnames(n int) (names []string, err error) {
	if m.ReaddirnamesFn == nil {
		return nil, ErrMockNotImplemented
	}
	return m.ReaddirnamesFn(n)
}

func (m *MockFile) ReadDir(n int) ([]os.DirEntry, error) {
	if m.ReadDirFn == nil {
		return nil, ErrMockNotImplemented
	}
	return m.ReadDirFn(n)
}

func (m *MockFile) Close() error {
	if m.CloseFn == nil {
		return ErrMockNotImplemented
	}
	return m.CloseFn()
}

func (m *MockFile) Chown(uid int, gid int) error {
	if m.ChownFn == nil {
		return ErrMockNotImplemented
	}
	return m.ChownFn(uid, gid)
}

func (m *MockFile) Truncate(size int64) error {
	if m.TruncateFn == nil {
		return ErrMockNotImplemented
	}
	return m.TruncateFn(size)
}

func (m *MockFile) Sync() error {
	if m.SyncFn == nil {
		return ErrMockNotImplemented
	}
	return m.SyncFn()
}

func (m *MockFile) Chdir() error {
	if m.ChdirFn == nil {
		return ErrMockNotImplemented
	}
	return m.ChdirFn()
}

func (m *MockFile) Stat() (os.FileInfo, error) {
	if m.StatFn == nil {
		return nil, ErrMockNotImplemented
	}
	return m.StatFn()
}

func (m *MockFile) Name() string {
	if m.NameFn == nil {
		return ""
	}
	return m.NameFn()
}

func (m *MockFile) Read(b []byte) (n int, err error) {
	if m.ReadFn == nil {
		return 0, ErrMockNotImplemented
	}
	return m.ReadFn(b)
}

func (m *MockFile) ReadAt(b []byte, off int64) (n int, err error) {
	if m.ReadAtFn == nil {
		return 0, ErrMockNotImplemented
	}
	return m.ReadAtFn(b, off)
}

func (m *MockFile) ReadFrom(r io.Reader) (n int64, err error) {
	if m.ReadFromFn == nil {
		return 0, ErrMockNotImplemented
	}
	return m.ReadFromFn(r)
}

func (m *MockFile) Write(b []byte) (n int, err error) {
	if m.WriteFn == nil {
		return 0, ErrMockNotImplemented
	}
	return m.WriteFn(b)
}

func (m *MockFile) WriteAt(b []byte, off int64) (n int, err error) {
	if m.WriteAtFn == nil {
		return 0, ErrMockNotImplemented
	}
	return m.WriteAtFn(b, off)
}

func (m *MockFile) Seek(offset int64, whence int) (ret int64, err error) {
	if m.SeekFn == nil {
		return 0, ErrMockNotImplemented
	}
	return m.SeekFn(offset, whence)
}

func (m *MockFile) WriteString(s string) (n int, err error) {
	if m.WriteStringFn == nil {
		return 0, ErrMockNotImplemented
	}
	return m.WriteStringFn(s)
}

func (m *MockFile) Chmod(mode os.FileMode) error {
	if m.ChmodFn == nil {
		return ErrMockNotImplemented
	}
	return m.ChmodFn(mode)
}

func (m *MockFile) SetDeadline(t time.Time) error {
	if m.SetDeadlineFn == nil {
		return ErrMockNotImplemented
	}
	return m.SetDeadlineFn(t)
}

func (m *MockFile) SetReadDeadline(t time.Time) error {
	if m.SetReadDeadlineFn == nil {
		return ErrMockNotImplemented
	}
	return m.SetReadDeadlineFn(t)
}

func (m *MockFile) SetWriteDeadline(t time.Time) error {
	if m.SetWriteDeadlineFn == nil {
		return ErrMockNotImplemented
	}
	return m.SetWriteDeadlineFn(t)
}

func (m *MockFile) SyscallConn() (syscall.RawConn, error) {
	if m.SyscallConnFn == nil {
		return nil, ErrMockNotImplemented
	}
	return m.SyscallConnFn()
}

func (m *MockFile) Fd() uintptr {
	if m.FdFn == nil {
		return 0
	}
	return m.FdFn()
}
