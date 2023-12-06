package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func main() {
	argLength := len(os.Args[1:])

	if argLength < 1 {
		panic("Not enough arguments")
	}

	day := os.Args[1]

	dirName := filepath.Join("./", day)

	fmt.Printf("Creating Day %s directory\n", dirName)

	_, err := os.Stat(dirName)

	if !os.IsNotExist(err) {
		fmt.Printf("%s Directory already exists\n", dirName)
		return
	}

	err = os.Mkdir(dirName, 0755)

	if err != nil {
		panic(err)
	}

	createMainFile(day)
	createInputFile(day)
	createPartFile(day, 1, "One")
	createPartFile(day, 2, "Two")
}

func createInputFile(day string) {
	filename := fmt.Sprintf("%s.txt", day)
	filePath := filepath.Join("input", filename)
	f, err := os.Create(filePath)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s created", filePath)

	f.Close()
}

func createMainFile(day string) {
	filePath := filepath.Join(day, "main.go")
	f, err := os.Create(filePath)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	fmt.Printf("%s created\n", filePath)

	templ := template.Must(template.New("").Parse(mainFileTeml))

	templ.Execute(f, map[string]string{"Day": day})
}

func createPartFile(day string, part int, partName string) {
	filename := fmt.Sprintf("part%d.go", part)
	filePath := filepath.Join(day, filename)

	f, err := os.Create(filePath)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s created\n", filePath)

	defer f.Close()

	data := map[string]interface{}{
		"Day":      day,
		"Part":     part,
		"PartName": partName,
	}

	templ := template.Must(template.New("").Parse(partFileTempl))
	templ.Execute(f, data)
}

var mainFileTeml = `
package main

import (
	"aoc/runner"
)

func main() {
	runner := runner.NewRunner("input/{{ .Day }}.txt")

	partone := NewPartOne()
	runner.Run(partone)

  // Uncomment to work on part two
	// parttwo := NewPartTwo()
	// runner.Run(parttwo)
}
`
var partFileTempl = `
package main

import (
	"fmt"
)

type Part{{ .PartName }} struct {
}

func NewPart{{ .PartName }}() *Part{{ .PartName }} {
  return &Part{{ .PartName }} {
  }
}

func (p *Part{{ .PartName }}) Init() {
	fmt.Println("Initializing Day {{ .Day }} Part {{ .Part }}")
}

func (p *Part{{ .PartName }}) ReadLine(line string, idx int) {
  fmt.Println("Reading line", line, idx)
}

func (p *Part{{ .PartName }}) Finish() {
	fmt.Printf("Day {{ .Day }} Part {{ .Part }}:\n")
	// fmt.Printf("%d:\n", p.Sum)
}
`
