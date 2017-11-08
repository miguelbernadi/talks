package main

// Reader interface (like io.Reader)
type Reader interface {
	Read(p []byte) (n int, err error)
}

// Valid struct (empty)
type Valid struct{}

// Read method on Valid so it fulfills Reader interface
func (v Valid) Read(p []byte) (n int, err error) {
}

// Wrong struct (empty)
type Wrong struct{}

// Read method on Wrong doesn't match Reader
func (w Wrong) Read(str string) (n int, err error) {
}
