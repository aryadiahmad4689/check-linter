package runner

import (
	"os"

	"github.com/aryadiahmad4689/check-linter/src/find"
	"github.com/aryadiahmad4689/check-linter/src/parser"
	_ "github.com/joho/godotenv/autoload"
)

type Runner struct {
}

type Antention struct {
	Data []parser.Data
}

func Init() *Runner {
	return &Runner{}
}
func (r *Runner) Excute() *Antention {
	var Antention = &Antention{}
	for _, s := range find.Find(os.Getenv("PATH_TO_LINTER"), ".go") {
		oke := parser.Init().Parser(s).Status
		if ok {
			Antention.Data = append(Antention.Data, parser.Init().Parser(s).Data)

		}
	}
	return Antention
}
