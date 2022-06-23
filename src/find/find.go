package find

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"

	config "github.com/aryadiahmad4689/check-linter/src/data"
)

func Find(root, ext string) []string {
	var a []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == ext {
			a = append(a, s)
		}
		return nil
	})
	return a
}

func FindConfig() string {
	var files string

	root := "./"

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".json" {
			if path == "check_linter.json" {
				files = path
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return files
}

func GetConfig() config.Config {
	data := FindConfig()
	jsonFile, err := os.Open(data)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println("check_linter.json tidak ditemukan")
		return config.Config{}
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var config = &config.Config{}
	json.Unmarshal([]byte(byteValue), config)

	return *config
}
