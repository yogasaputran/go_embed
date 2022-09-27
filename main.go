package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed version.txt
var version string

//go:embed logo.jpeg
var logo []byte

//go:embed files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)
	err := ioutil.WriteFile("logo_new_1.jpeg", logo, fs.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}
