package wc

import (
	"bufio"
	"io"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"unicode/utf8"
)

type Counter struct {
	f *os.File

	Multibytes uint64
	Bytes      uint64
	Lines      uint64
	Words      uint64
}

var BufferSize = 4 << 20

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func NewCounter(f *os.File) *Counter {
	return &Counter{
		f: f,
	}
}

func (c *Counter) Count(count_multibytes, count_bytes, count_lines, count_words bool) error {

	workers, wg := runtime.NumCPU()*2, new(sync.WaitGroup)

	wg.Add(workers)

	buffers := make(chan []byte, workers<<3)

	for i := 0; i < workers; i++ {
		go func() {

			var multibyte_ct, byte_ct, line_ct, word_ct uint64

			for buf := range buffers {

				if count_multibytes {
					multibyte_ct += uint64(utf8.RuneCount(buf))
				}
				if count_bytes {
					byte_ct += uint64(len(buf))
				}

				if count_words {
					var in_word bool
					for _, c := range buf {
						if isSpace(c) {
							if c == '\n' {
								line_ct++
							}
							in_word = false
						} else if !in_word {
							word_ct++
							in_word = true
						}
					}
				} else if count_lines {
					for _, c := range buf {
						if c == '\n' {
							line_ct++
						}
					}
				}

			}

			atomic.AddUint64(&c.Multibytes, multibyte_ct)
			atomic.AddUint64(&c.Lines, line_ct)
			atomic.AddUint64(&c.Words, word_ct)
			atomic.AddUint64(&c.Bytes, byte_ct)

			wg.Done()
		}()
	}

	br := bufio.NewReaderSize(c.f, BufferSize)

	for {
		buf := make([]byte, BufferSize)

		n, err := br.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		buf = buf[:n]

		// read to a newline to prevent
		// spanning a word across two buffers
		if count_words && buf[len(buf)-1] != '\n' {
			extra, err := br.ReadBytes('\n')
			if err != nil && err != io.EOF {
				return err
			}
			buf = append(buf, extra...)
		}

		buffers <- buf
	}

	close(buffers)
	wg.Wait()

	return nil

}

// isspace() man page spec
// http://linux.die.net/man/3/isspace
func isSpace(c byte) bool {
	return c == ' ' || c == '\n' || c == '\r' || c == '\t' || c == '\v' || c == '\f'
}
