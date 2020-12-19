package statJanitor

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func TestCtxArg(t *testing.T) {
	analysistest.Run(t, analysistest.TestData() + "/Internet2", Analyzer)
}
