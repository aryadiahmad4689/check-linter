package visitor

import (
	"fmt"
	"go/ast"
	"go/token"
)

type Visitor struct {
	Fset *token.FileSet
	Data []string
}

func (v *Visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}
	switch x := n.(type) {
	case *ast.ExprStmt:
		fmst, ok := x.X.(*ast.CallExpr)

		if ok {
			dta, oks := fmst.Fun.(*ast.SelectorExpr)
			if oks {
				if dta.Sel.Name == "Print" && dta.X.(*ast.Ident).Name == "fmt" {
					data := fmt.Sprintf("Ditemukan fmt.Print di line dan file:%s", v.Fset.Position(n.Pos()))
					v.Data = append(v.Data, data)
				} else if dta.Sel.Name == "Println" && dta.X.(*ast.Ident).Name == "fmt" {
					data := fmt.Sprintf("Ditemukan fmt.Println line dan file:%s", v.Fset.Position(n.Pos()))
					v.Data = append(v.Data, data)
				}
			}
		}

	case *ast.AssignStmt:
		if x.Tok.String() == ":=" {
			data := fmt.Sprintf("Ditemukan := di line dan file:%s", v.Fset.Position(n.Pos()))
			v.Data = append(v.Data, data)
		}

	}

	return v
}
