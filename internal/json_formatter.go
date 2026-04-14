// JSONFormatter outputs an envelope containing a summary object and a
// results array. Every checked link appears in the results — not just
// broken ones — because consumers often need to know what was checked,
// not just what failed. The schema is kept uniform: "reason" is always
// present as a string (empty when the link is valid) so consumers do
// not need null-handling logic.
//
// The output is indented for human readability. Machine consumers are
// unaffected since JSON parsers ignore whitespace.
package internal

import (
	"encoding/json"
	"io"
)

// JSONFormatter writes JSON envelope output with summary and results.
type JSONFormatter struct{}

// jsonResult is the per-link structure in the JSON output. It flattens
// the Link fields to avoid nested objects — the consumer does not need
// to know about mdcheck's internal Link type.
type jsonResult struct {
	File   string `json:"file"`
	Line   int    `json:"line"`
	Text   string `json:"text"`
	Target string `json:"target"`
	Broken bool   `json:"broken"`
	Reason string `json:"reason"`
}

// jsonEnvelope is the top-level JSON structure.
type jsonEnvelope struct {
	Summary struct {
		Total  int `json:"total"`
		Broken int `json:"broken"`
	} `json:"summary"`
	Results []jsonResult `json:"results"`
}

// Format writes the JSON envelope to w.
func (f JSONFormatter) Format(results []Result, w io.Writer) error {
	broken := 0
	jsonResults := make([]jsonResult, 0, len(results))

	for _, r := range results {
		if r.Broken {
			broken++
		}
		jsonResults = append(jsonResults, jsonResult{
			File:   r.Link.File,
			Line:   r.Link.Line,
			Text:   r.Link.Text,
			Target: r.Link.Target,
			Broken: r.Broken,
			Reason: r.Reason,
		})
	}

	env := jsonEnvelope{
		Results: jsonResults,
	}
	env.Summary.Total = len(results)
	env.Summary.Broken = broken

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(env)
}
