package errgroupctx_test

import (
	"testing"

	"github.com/ndx-technologies/lint-errgroup-ctx/analysis/errgroupctx"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), errgroupctx.Analyzer, "./example")
}
