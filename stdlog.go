// Package stdlog provides a go/analysis analyzer enforcing the gomatic Go logging
// standard: the standard library "log" package is forbidden in favor of
// structured logging with "log/slog".
package stdlog

import (
	goyze "github.com/gomatic/go-yze"
	"golang.org/x/tools/go/analysis"
)

const (
	message = "the standard log package is forbidden; use log/slog for structured logging"
	// stdLogPath is the quoted import path of the standard log package, as it
	// appears in an ast.ImportSpec's Path.Value.
	stdLogPath = `"log"`
)

// Analyzer reports imports of the standard "log" package.
var Analyzer = &analysis.Analyzer{
	Name: "stdlog",
	Doc:  "reports imports of the standard log package, which the gomatic standard replaces with log/slog",
	Run:  run,
}

// Registration declares this analyzer to the yze framework.
var Registration = goyze.Registration{
	Name:       "stdlog",
	Categories: []goyze.Category{"data"},
	URL:        "https://docs.gomatic.dev/yze/stdlog",
	Analyzer:   Analyzer,
}

// run reports each import of the standard log package.
func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		for _, imp := range file.Imports {
			if imp.Path.Value == stdLogPath {
				pass.Reportf(imp.Pos(), message)
			}
		}
	}
	return nil, nil
}
