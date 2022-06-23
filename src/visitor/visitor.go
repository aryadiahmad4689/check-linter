package visitor

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/aryadiahmad4689/check-linter/src/data"
)

type Visitor struct {
	Fset *token.FileSet
	Data []string
	Conf data.Config
}

func (v *Visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}
	conf := v.Conf
	switch x := n.(type) {
	case *ast.ExprStmt:
		fmst, ok := x.X.(*ast.CallExpr)

		if ok {
			dta, oks := fmst.Fun.(*ast.SelectorExpr)
			if oks {
				for _, conf := range conf.SelectorExpression {
					if dta.Sel.Name == conf.Sel && dta.X.(*ast.Ident).Name == conf.X {
						data := fmt.Sprintf("Ditemukan %s.%s di line dan file:%s", conf.X, conf.Sel, v.Fset.Position(n.Pos()))
						v.Data = append(v.Data, data)
					}
				}

			}
		}

	case *ast.AssignStmt:
		for _, conf := range conf.ExprStmt {
			if x.Tok.String() == conf.Sel {
				data := fmt.Sprintf("Ditemukan %s di line dan file:%s", conf.Sel, v.Fset.Position(n.Pos()))
				v.Data = append(v.Data, data)
			}
		}

	case *ast.DeclStmt:
		for _, conf := range conf.ExprStmt {
			if x.Decl.(*ast.GenDecl).Tok.String() == conf.Sel {
				data := fmt.Sprintf("Ditemukan %s di line dan file:%s", conf.Sel, v.Fset.Position(n.Pos()))
				v.Data = append(v.Data, data)
			}
		}
	}

	return v
}
