package main

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"text/template"
)

const part1 = `package {{.Day}}

import (
	"fmt"

	"github.com/gomsim/Advent-of-code/shared/register"	
)

func init() {
	register.Part1("2024", "{{.Day}}", Part1)
}
	
func Part1(file string) {
	fmt.Println("NOT IMPLEMENTED!")
}`

const part2 = `package {{.Day}}

import "github.com/gomsim/Advent-of-code/shared/register"

func init() {
	register.Part2("2024", "{{.Day}}", Part2)
}
	
func Part2(file string) {
	fmt.Println("NOT IMPLEMENTED!")
}`

func main() {
	day := os.Args[1]

	solution := map[string]map[string]string{
		day: {
			"part1.go": part1,
			"part2.go": part2,
		},
	}

	input := map[string]map[string]string{
		day: {
			"input.txt": "",
			"test.txt":  "",
		},
	}

	data := map[string]any{
		"Day": day,
	}

	create("solutions", solution, data)
	create("input", input, data)

	fmt.Println("Catalogue and files with dynamic content created successfully.")

	addImport(day)
}

func create(parentDir string, structure map[string]map[string]string, data map[string]any) {
	for dir, files := range structure {
		err := os.MkdirAll(path.Join("2024", parentDir, dir), 0755)
		if err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dir, err)
			continue
		}

		for fileName, templateContent := range files {
			filePath := path.Join("2024", parentDir, dir, fileName)

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

			data["FileName"] = fileName

			err = tmpl.Execute(file, data)
			if err != nil {
				fmt.Printf("Error writing to file %s: %v\n", filePath, err)
			}

			file.Close()
		}
	}
}

func addImport(day string) {
	mainFilePath := "main.go"
	importStatement := fmt.Sprintf(`_ "github.com/gomsim/Advent-of-code/2024/solutions/%s"`, day)

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
