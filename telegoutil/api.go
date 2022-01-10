package telegoutil

import (
	"io"

	"github.com/mymmrac/telego/telegoapi"
)

// namedReaderImpl represents simplest implementation of telegoapi.NamedReader
type namedReaderImpl struct {
	reader io.Reader
	name   string
}

func (r namedReaderImpl) Read(p []byte) (n int, err error) {
	return r.reader.Read(p)
}

func (r namedReaderImpl) Name() string {
	return r.name
}

// NameReader "names" io.Reader and returns valid telegoapi.NamedReader
func NameReader(reader io.Reader, name string) telegoapi.NamedReader {
	return namedReaderImpl{
		reader: reader,
		name:   name,
	}
}
