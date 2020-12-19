package main

import (
	"github.com/amit-davidson/HttpBodyClose"
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func TestCtxArg(t *testing.T) {
	analysistest.Run(t, analysistest.TestData() + "/Internet2", HttpBodyClose.Analyzer)
}
