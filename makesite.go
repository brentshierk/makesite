package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)



func main() {
	var fileName,directory string

	flag.StringVar(&fileName,"file","","text file you want to turn into an html page")
	flag.StringVar(&directory,"dir",".","directory to look for text files to convert (all)")
	flag.Parse()

	

	if directory != ""{
		dirToHtml(directory)
	}else if fileName != ""{
		fileToHtml(fileName)

	}

	
}

func fileToHtml(fileName string){
	fileContent, err := ioutil.ReadFile(fileName)
	checkErr(err)

	f,err := os.Create(strings.SplitN(fileName,".",2)[0]+".html")
	checkErr(err)

	t:= template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	err = t.Execute(f,string(fileContent))
	checkErr(err)

	f.Close()
}
func dirToHtml(directory string){
	files,err := ioutil.ReadDir(directory)
	checkErr(err)

	for _,file := range files{
		if file.Name()[len(file.Name())-4:]==".txt"{
			fileToHtml(file.Name())
		}

	}

}


func checkErr(err error){
	if err != nil{
		println(err)
	}
}


