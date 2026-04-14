// The Formatter interface decouples output rendering from link checking.
// Each implementation writes results in a specific format (text, JSON,
// GitHub annotations) to the provided writer. This separation means the
// checker pipeline does not need to know how results are displayed, and
// new formats can be added by implementing a single method.
//
// The interface takes an io.Writer rather than returning a string so that
// formatters can stream output and callers control the destination.
package internal

import "io"

// Formatter renders check results to a writer in a specific output format.
type Formatter interface {
	Format(results []Result, w io.Writer) error
}
