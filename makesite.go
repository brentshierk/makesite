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
	flag.StringVar(&directory,"dir","","direcotry to look for text files to convert (all)")
	flag.Parse()

	if TextFilePath == "" && directory ==""{
		panic("missing the --flag or --dir flag please suply one or the other")

	}

	if directory != ""{
		createallFiles(directory)
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
func createallFiles(directory string,)(page Page){
	
	files,err := os.ReadDir(directory)

	if err != nil{
		print(err)
	}

	for _,file := range files{
		path := directory + "/" +file.Name()

		info,err := os.Stat(path)
		if err !=nil{
			println(err)
		}

		if info.IsDir(){
			createallFiles(path)
		}else {
			if isTxt(file.Name()) {
				fileContent := readFile(path)
				page:= createPageFromTextFile("template.tmpl" + path+string(fileContent))
				println(page)

			}
		}
	}

	return page
}
func readFile(file string) []byte {
	content, err := os.ReadFile(file)
	if err !=nil{
		println(err)
	}
	return content
}
func isTxt(filename string) bool {
	fileExt := filename[len(filename)-4:]
	return fileExt == ".txt"
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


