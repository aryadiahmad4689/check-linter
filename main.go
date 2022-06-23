package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aryadiahmad4689/check-linter/src/runner"

	_ "github.com/joho/godotenv/autoload"
)

func CheckCode() runner.Antention {
	var data = runner.Init(runner.InitR{
		Path: "src",
	}).Excute()
	return *data
}

func main() {
	var rootDir = flag.String("root_dir", "", "Ini adalah flag root directory exmaples: -root_dir=src/halo")
	flag.Parse()
	if *rootDir == "" {
		fmt.Println("root_dir harus di isi")
		os.Exit(1)
	}
	runner.Init(runner.InitR{
		Path: *rootDir,
	}).Excute()

	// fmt.Println(config.SelectorExpression)
}
