# Resumes

Each resume has its own subdirectory and data.yml file. The resume is generated using the tmpl directory. You may find it useful to recreate this setup for your own resumes.

## Disclaimer

Please do not run scripts from sources that you do not trust. Do take time to verify and understand how the script works before running it.

## Resources for Learning Go

The resources that helped in understanding how the script works. They are not listed in any order.

* [A Tour of Go](https://tour.golang.org/)
* [Go by Example](https://gobyexample.com/)
* [Head First Go](https://headfirstgo.com/)
* [Practical Go Lessons](https://www.practical-go-lessons.com/)
* [Astaxie's GitBook](https://astaxie.gitbooks.io/build-web-application-with-golang/content/)

## Instructions for Resume

This resume is built from YAML structured data (which is
drawn from the CV) and generated from [templates](tmpl) in Go
(golang) `html/template` syntax. The PDF is created with Firefox.

1. Create a new subdirectory and `data.yml` inside the subdirectory.
1. Write the resume in `data.yml`
1. Generate the `index.html` file from `tmpl` through `go run gen.go <source directory>` from within this
directory.
1. Open the `index.html` in Firefox.
1. Create the PDF by printing the page generated to save as a PDF.