# fi - File Interface
An interface exposing all exported methods of *os.File.

## Rationale 
In an internal API that makes heavy use of file operations, it can be helpful
to mock the behaviour of a file for testing, or inject functionality for
debugging.

In particular, testing code to be resilient against various infrequent OS
errors can be difficult because these errors can't be easily simulated.

## Usage

Godoc: http://pkg.go.dev/github.com/iamcalledrob/fi

### Interface
Use `fi.File` in place of `*os.File` in internal APIs.
```go
func wrangle(f fi.File) error {
	// ...
}
```

### MockFile
`fi.MockFile` implements fi.File, and allows for a fully mocked file 
```go
f := &fi.MockFile{
	ReadFn: func(p []byte) (int, error) {
		// ...
	},
	WriteFn: func(b []byte) (int, error) {
		// ...
	}
}

wrangle(f)
```

### MockFileFromOS
`fi.MockFileFromOS()` delegates all functionality to an underlying
`*os.File`. Behaviour can then be tweaked on a per-method basis.

For example, to test write/close behaviour when the disk is out of space
```go
// raw is *os.File
raw, err := os.Open("/path/to/file")
if err != nil {
	// ...
}

f := fi.MockFileFromOS(raw)

f.CloseFn = func() error {
	return &os.PathError{
		Op:   "close",
		Path: f.Name(),
		Err:  syscall.ENOSPC,
	}
}

f.Write([]byte{1, 2, 3})
f.Close() // "out of space"
```
