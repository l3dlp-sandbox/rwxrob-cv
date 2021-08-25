package main

import (
	"html/template"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var templates *template.Template

func main() {
	source := os.Args[1]
	if source == "" {
		log.Fatal("usage go run gen.go <source directory>")
	}
	data := map[string]interface{}{}
	buf, err := os.ReadFile(source + "/data.yml")
	checkError(err)

	err = yaml.Unmarshal(buf, &data)
	checkError(err)

	var allFiles []string
	files, err := os.ReadDir("tmpl")
	checkError(err)

	for _, file := range files {
		fileName := file.Name()
		if strings.HasSuffix(fileName, ".html") {
			allFiles = append(allFiles, "tmpl/"+fileName)
		}

	}

	templates, err = template.ParseFiles(allFiles...)
	checkError(err)

	out, err := os.Create(source + "/index.html")
	defer out.Close()
	checkError(err)

	main := templates.Lookup("tmpl/_.html")
	main.Execute(out, data)

	// TODO detect tmpl and fail if not found
	// TODO iterate over tmpl directory and
	//    if directory build a file matching the name of the directory
	//    from the files in the directory
	//    or,
	//    if a file just build from that file
	//    make sure to detect the template/html or template/text based on
	//    suffix
}

// */
