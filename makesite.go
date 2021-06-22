package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

type Page struct{
	TextFilePath string
	TextFileName string
	HTMLPagePath string
	Content string
}

func main() {
	var TextFilePath,directory string
	flag.StringVar(&TextFilePath,"file","","text file you want to turn into an html page")
	flag.StringVar(&directory,"dir",".","direcotry to look for text files to convert (all)")
	flag.Parse()

	if TextFilePath == "" && directory ==""{
		panic("missing the --flag or --dir flag please suply one or the other")

	}

	newPage := createPageFromTextFile(TextFilePath)

	renderTemplateFromPage("template.tmpl",newPage)

}

func renderTemplateFromPage(templateFilePath string,page Page){
	t:= template.Must(template.New(templateFilePath).ParseFiles(templateFilePath))

	newFile,err := os.Create(page.HTMLPagePath)
	if err != nil{
		panic(err)

	}

	t.Execute(newFile,page)
	fmt.Println("generated file :",page.HTMLPagePath)
}

func createPageFromTextFile(filePath string)Page{
	fileContents , err := ioutil.ReadFile(filePath)
	if err != nil{
		panic(err)
	}

	fileNameWithoutExtension := strings.Split(filePath,".txt")[0]

	return Page{
		TextFilePath: filePath,
		TextFileName: fileNameWithoutExtension,
		HTMLPagePath: fileNameWithoutExtension + ".html",
		Content: string(fileContents),
	}

}
