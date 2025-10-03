package main

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"text/template"

	"github.com/alexflint/go-arg"
)

const part1 = `package {{.Day}}

import (
	"github.com/gomsim/Advent-of-code/shared/registrar"	
)

func init() {
	registrar.Register("{{.Year}}", "{{.Day}}", "part1", Part1)
}
	
func Part1(file string) any {
	//in := input.Slice(file)

	return "NOT IMPLEMENTED!"
}`

const part2 = `package {{.Day}}

import (
	"github.com/gomsim/Advent-of-code/shared/registrar"	
)

func init() {
	registrar.Register("{{.Year}}", "{{.Day}}", "part2", Part2)
}
	
func Part2(file string) any {
	//in := input.Slice(file)

	return "NOT IMPLEMENTED!"
}`

const common = `package {{.Day}}`

type input struct {
	Year string `arg:"-y,--year" default:"2024"`
	Day  string `arg:"-d,--day,required"`
}

func main() {
	in := input{}
	arg.MustParse(&in)
	in.Day = fmt.Sprintf("day%s", in.Day)

	solution := map[string]map[string]string{
		in.Day: {
			"part1.go":  part1,
			"part2.go":  part2,
			"common.go": common,
		},
	}

	input := map[string]map[string]string{
		in.Day: {
			"input.txt": "",
			"test.txt":  "",
		},
	}

	data := map[string]any{
		"Day":  in.Day,
		"Year": in.Year,
	}

	create("solutions", in.Year, solution, data)
	create("input", in.Year, input, data)

	fmt.Println("Catalogue and files with dynamic content created successfully.")

	addImport(in.Year, in.Day)
}

func create(parentDir string, year string, structure map[string]map[string]string, data map[string]any) {
	for dir, files := range structure {
		err := os.MkdirAll(path.Join(year, parentDir, dir), 0755)
		if err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dir, err)
			continue
		}

		for fileName, templateContent := range files {
			filePath := path.Join(year, parentDir, dir, fileName)

			tmpl, err := template.New(fileName).Parse(templateContent)
			if err != nil {
				fmt.Printf("Error parsing template for file %s: %v\n", filePath, err)
				continue
			}

			file, err := os.Create(filePath)
			if err != nil {
				fmt.Printf("Error creating file %s: %v\n", filePath, err)
				continue
			}

			err = tmpl.Execute(file, data)
			if err != nil {
				fmt.Printf("Error writing to file %s: %v\n", filePath, err)
			}

			file.Close()
		}
	}
}

func addImport(year string, day string) {
	mainFilePath := "main.go"
	importStatement := fmt.Sprintf(`_ "github.com/gomsim/Advent-of-code/%s/solutions/%s"`, year, day)

	content, err := os.ReadFile(mainFilePath)
	if err != nil {
		fmt.Printf("Error reading main.go: %v\n", err)
		return
	}

	if bytes.Contains(content, []byte(importStatement)) {
		fmt.Println("Import statement already exists in main.go")
		return
	}

	updatedContent := bytes.Replace(content, []byte("import ("), []byte(fmt.Sprintf("import (\n\t%s", importStatement)), 1)
	if bytes.Equal(updatedContent, content) {
		updatedContent = bytes.Replace(content, []byte("import"), []byte(fmt.Sprintf("import (\n\t%s\n)", importStatement)), 1)
	}

	err = os.WriteFile(mainFilePath, updatedContent, 0644)
	if err != nil {
		fmt.Printf("Error writing to main.go: %v\n", err)
		return
	}

	fmt.Println("Import statement added to main.go")
}
