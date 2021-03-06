package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"

	"github.com/aryadiahmad4689/check-linter/src/find"
	"github.com/aryadiahmad4689/check-linter/src/visitor"
)

type Data struct {
	Status bool
	Data   []string
}
type Parser struct {
}

func Init() *Parser {
	return &Parser{}
}
func (p *Parser) Parser(path string) Data {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		log.Fatal(err)
	}

	conf := find.GetConfig()

	visitor := &visitor.Visitor{Fset: fset, Conf: conf}
	ast.Walk(visitor, file)

	return Data{
		Status: len(visitor.Data) > 0,
		Data:   visitor.Data,
	}
}
