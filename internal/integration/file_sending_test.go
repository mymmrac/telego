//go:build integration && fileSending

package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func TestFileSending(t *testing.T) {
	ctx := t.Context()

	t.Log("Running pprof at: http://localhost:8080/debug/pprof/")
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			panic(err)
		}
	}()

	sendingCtx, cancel := context.WithCancel(ctx)
	go trackMemory(t, sendingCtx, time.Millisecond*500)

	time.Sleep(time.Millisecond * 100)

	msg, err := bot.SendDocument(ctx, &telego.SendDocumentParams{
		ChatID: tu.ID(chatID),
		Document: tu.File(&randomReader{
			length: 1024 * 1024 * 20,
			delay:  time.Millisecond * 20,
		}),
		Caption: "SendDocument " + timeNow,
	})

	require.NoError(t, err)
	assert.NotNil(t, msg)

	cancel()
	t.Log("Sending done")
	<-t.Context().Done()
}

type randomReader struct {
	length int
	delay  time.Duration
}

func (r *randomReader) Name() string {
	return "random.txt"
}

func (r *randomReader) Read(p []byte) (int, error) {
	n := len(p)
	if n > r.length {
		n = r.length
	}
	if n <= 0 {
		return 0, io.EOF
	}
	r.length -= n
	fmt.Printf("Left to read %d bytes\n", r.length)
	time.Sleep(r.delay)
	return rand.Read(p[:n])
}

func trackMemory(t *testing.T, ctx context.Context, interval time.Duration) {
	t.Helper()

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	var maxHeap uint64
	for {
		select {
		case <-ticker.C:
			var m runtime.MemStats
			runtime.ReadMemStats(&m)

			if m.HeapAlloc > maxHeap {
				maxHeap = m.HeapAlloc
			}

			t.Logf("HeapAlloc = %d MB\n", m.HeapAlloc/1024/1024)
			t.Logf("HeapSys   = %d MB\n", m.HeapSys/1024/1024)
			t.Logf("Sys       = %d MB\n", m.Sys/1024/1024)
		case <-ctx.Done():
			t.Logf("Peak HeapAlloc: %.2f MB\n", float64(maxHeap)/1024/1024)
			return
		}
	}
}
