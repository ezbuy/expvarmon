package fetch

import (
	"io"

	"github.com/antonholmquist/jason"
)

// Expvar represents fetched expvar variable.
type Expvar struct {
	*jason.Object
}

// ParseExpvar parses expvar data from reader.
func ParseExpvar(r io.Reader) (*Expvar, error) {
	object, err := jason.NewObjectFromReader(r)
	return &Expvar{object}, err
}
