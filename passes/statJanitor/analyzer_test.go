package statJanitor

import (
	"fmt"
	"golang.org/x/tools/go/analysis/analysistest"
	"os"
	"testing"
)

func TestCtxArg(t *testing.T) {
	var testCases = []struct {
		name     string
		testPath string
	}{
		{name: "Internet1"},
		{name: "Internet2"},
		{name: "passBodyFail"},
		{name: "passBodySuccess"},
		{name: "passCloseFail"},
		{name: "passCloseSuccess"},
		{name: "passHttpResponseFail"},
		{name: "passHttpResponseSuccess"},
	}
	for _, tc := range testCases {
		analysistest.Run(t, fmt.Sprintf("%s%s%s", analysistest.TestData(), string(os.PathSeparator), tc.name), Analyzer)
	}
}
