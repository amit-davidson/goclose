package main

import (
	"github.com/amit-davidson/HttpBodyClose"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(HttpBodyClose.Analyzer)
}
