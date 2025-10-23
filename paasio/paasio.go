package paasio

import (
	"io"
	"sync"
)

// counter is a shared delegate for counting bytes and ops.
// We don't use atomic variables because those work on only one variable at a time.
// If we need to protect updates to multiple variables together, we need to use mutexes
// or other synchronization tools.
type counter struct {
	mu    sync.Mutex
	bytes int64
	ops   int
}

// add updates the counter safely.
func (c *counter) add(n int) {
	c.mu.Lock()
	c.bytes += int64(n)
	c.ops++
	c.mu.Unlock()
}

// snapshot returns a consistent snapshot of the counter.
func (c *counter) snapshot() (int64, int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.bytes, c.ops
}

// readCounter implements ReadCounter using a delegate counter.
type readCounter struct {
	r io.Reader
	c *counter
}

func NewReadCounter(r io.Reader) ReadCounter {
	return &readCounter{r: r, c: &counter{}}
}

func (rc *readCounter) Read(p []byte) (int, error) {
	n, err := rc.r.Read(p)
	rc.c.add(n)
	return n, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	return rc.c.snapshot()
}

// writeCounter implements WriteCounter using a delegate counter.
type writeCounter struct {
	w io.Writer
	c *counter
}

func NewWriteCounter(w io.Writer) WriteCounter {
	return &writeCounter{w: w, c: &counter{}}
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	n, err := wc.w.Write(p)
	wc.c.add(n)
	return n, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	return wc.c.snapshot()
}

type readWriteCounter struct {
	readCounter
	writeCounter
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{
		readCounter:  readCounter{r: rw, c: &counter{}},
		writeCounter: writeCounter{w: rw, c: &counter{}},
	}
}

// Methods just delegate
func (rwc *readWriteCounter) Read(p []byte) (int, error) {
	return rwc.readCounter.Read(p)
}

func (rwc *readWriteCounter) ReadCount() (int64, int) {
	return rwc.readCounter.ReadCount()
}

func (rwc *readWriteCounter) Write(p []byte) (int, error) {
	return rwc.writeCounter.Write(p)
}

func (rwc *readWriteCounter) WriteCount() (int64, int) {
	return rwc.writeCounter.WriteCount()
}
