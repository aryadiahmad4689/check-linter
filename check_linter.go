package check_linter

import (
	"github.com/aryadiahmad4689/check-linter/src/runner"
	_ "github.com/joho/godotenv/autoload"
)

func Parser() runner.Antention {
	var data = runner.Init().Excute()
	return *data
}
