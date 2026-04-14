// TextFormatter reproduces the original mdcheck output: one line per broken
// link showing file, line number, link text, target, and reason, followed
// by a summary line. Valid links produce no output lines — only the summary
// mentions them via the total count.
//
// This formatter exists so that the default behaviour is preserved exactly
// when users add --format support. If you change the output here, you change
// what every existing user sees.
package internal

import (
	"fmt"
	"io"
)

// TextFormatter writes human-readable plain text output.
type TextFormatter struct{}

// Format writes broken link details and a summary line to w.
func (f TextFormatter) Format(results []Result, w io.Writer) error {
	broken := 0
	for _, r := range results {
		if r.Broken {
			broken++
			_, err := fmt.Fprintf(w, "%s:%d: [%s](%s) — %s\n",
				r.Link.File, r.Link.Line, r.Link.Text, r.Link.Target, r.Reason)
			if err != nil {
				return err
			}
		}
	}

	_, err := fmt.Fprintf(w, "\n%d links checked, %d broken\n", len(results), broken)
	return err
}
