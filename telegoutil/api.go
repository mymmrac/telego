package telegoutil

import (
	"io"

	"github.com/mymmrac/telego"
	ta "github.com/mymmrac/telego/telegoapi"
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
func NameReader(reader io.Reader, name string) ta.NamedReader {
	return namedReaderImpl{
		reader: reader,
		name:   name,
	}
}

// UpdateProcessor allows you to process updates and still use updates chan. New updates chan will be closed when
// original chan is closed.
// Warning: Deep copy of update is passed, telego.Update's Clone() method can panic, please read its comment.
func UpdateProcessor(updates <-chan telego.Update, buffer uint, processor func(update telego.Update) telego.Update,
) <-chan telego.Update {
	processedUpdates := make(chan telego.Update, buffer)

	go func() {
		defer close(processedUpdates)
		for update := range updates {
			processedUpdates <- processor(update.Clone())
		}
	}()

	return processedUpdates
}
