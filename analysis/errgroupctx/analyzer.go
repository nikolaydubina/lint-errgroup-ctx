package errgroupctx

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "errgroupctx",
	Doc:      "checks that context from errgroup.WithContext is not assigned to 'ctx'",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (any, error) {
	ins := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	ins.Preorder([]ast.Node{(*ast.AssignStmt)(nil)}, func(n ast.Node) {
		assign := n.(*ast.AssignStmt)
		if assign == nil || len(assign.Lhs) != 2 || len(assign.Rhs) != 1 {
			return
		}
		call, ok := assign.Rhs[0].(*ast.CallExpr)
		if !ok || call == nil {
			return
		}
		sel, ok := call.Fun.(*ast.SelectorExpr)
		if !ok || sel == nil {
			return
		}
		obj := pass.TypesInfo.ObjectOf(sel.Sel)
		if obj == nil {
			return
		}
		funcObj, ok := obj.(*types.Func)
		if !ok {
			return
		}
		if funcObj.FullName() != "golang.org/x/sync/errgroup.WithContext" {
			return
		}
		ident, ok := assign.Lhs[1].(*ast.Ident)
		if !ok || ident.Name != "ctx" {
			return
		}
		pass.Reportf(ident.Pos(), "context from errgroup.WithContext should not be named 'ctx'")
	})

	return nil, nil
}
