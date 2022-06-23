package data

type SelectorExpr struct {
	Sel string
	X   string
}

type ExprStmt struct {
	Sel string
}

type DecStmt struct {
	Sel string
}

type Config struct {
	SelectorExpression []SelectorExpr
	ExprStmt           []ExprStmt
	DecStmt            []DecStmt
}
