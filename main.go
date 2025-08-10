package main

import (
	"github.com/nikolaydubina/lint-errgroup-ctx/analysis/errgroupctx"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(errgroupctx.Analyzer) }
