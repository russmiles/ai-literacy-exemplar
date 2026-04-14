// GitHubFormatter produces GitHub Actions workflow command annotations.
// GitHub's runner parses lines matching ::error file=F,line=L::message
// and renders them as inline annotations on pull request diffs. This
// makes broken links visible exactly where they occur in the source.
//
// Only broken links produce ::error annotations — valid links generate
// no output because annotations are a problem-reporting mechanism, not
// a status mechanism. The summary is emitted as a ::notice so it
// appears in the Actions log without cluttering the PR diff.
package internal

import (
	"fmt"
	"io"
)

// GitHubFormatter writes GitHub Actions annotation output.
type GitHubFormatter struct{}

// Format writes ::error annotations for broken links and a ::notice summary.
func (f GitHubFormatter) Format(results []Result, w io.Writer) error {
	broken := 0
	for _, r := range results {
		if r.Broken {
			broken++
			_, err := fmt.Fprintf(w, "::error file=%s,line=%d::Broken link: [%s](%s) — %s\n",
				r.Link.File, r.Link.Line, r.Link.Text, r.Link.Target, r.Reason)
			if err != nil {
				return err
			}
		}
	}

	_, err := fmt.Fprintf(w, "::notice::%d links checked, %d broken\n", len(results), broken)
	return err
}
