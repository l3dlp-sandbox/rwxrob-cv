# Generating Your Resume

You may find it useful to recreate this setup for your own resumes.

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

This resume is built from [YAML structured data](data.yml) (which is
drawn from the CV) and generated from [templates](tmpl) in Go
(golang) `html/template` syntax. To generate the `index.html` file from
updated `tmpl` simply do `go run gen.go` from within this
directory.

The PDF is created with Firefox with all the extras disabled and has
been tested to ensure it contains transferable fonts. Create the PDF by printing the page generated to save as a PDF.