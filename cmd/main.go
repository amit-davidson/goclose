package main

import (
	"github.com/amit-davidson/statJanitor/passes/statJanitor"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(statJanitor.Analyzer)
}
