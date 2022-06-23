package runner

import (
	"fmt"

	"github.com/aryadiahmad4689/check-linter/src/find"
	"github.com/aryadiahmad4689/check-linter/src/parser"
	_ "github.com/joho/godotenv/autoload"
)

type Runner struct {
	InitR
}

type Antention struct {
	Data []string
}

type InitR struct {
	Path string
}

func Init(initr InitR) *Runner {
	return &Runner{
		InitR: initr,
	}
}
func (r *Runner) Excute() *Antention {
	var Antention = &Antention{}
	for _, s := range find.Find(r.Path, ".go") {
		result := parser.Init().Parser(s)
		if result.Status {
			Antention.Data = append(Antention.Data, result.Data...)
		}
	}
	for _, result := range Antention.Data {
		fmt.Println(result)
	}
	return Antention
}
