package main

import (
	"flag"
	//"fmt"
	"html/template"
	"io/ioutil"
	"os"
	//"path/filepath"
	"strings"
	"spotify"
	"http"
)



func main() {
	var fileName,directory string

	flag.StringVar(&fileName,"file","","text file you want to turn into an html page")
	flag.StringVar(&directory,"directory",".","directory to look for text files to convert (all)")
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

func makePlaylist(file string){
ClientID := "32d484634e5843ee80024a0d65d1459d"
ClientSecret := "4629c73660b8499fa4809255e4bba77e"
auth.SetAuthInfo(clientID, secretKey)

//request_url := "https://api.spotify.com/v1/search?q="+track_name+"&type=track&market=US&limit=1&offset=5"
//var track_name string 

//read in the file 
fileContent, err := ioutil.ReadFile(fileName)
	checkErr(err)
//split the content of the file at ","
// for each substring compleate a query to the spotify api
//return json of playlist
//get  playlist icon,title,link to playlist



results, err := client.Search(song string , spotify.SearchTypePlaylist|spotify.SearchTypeAlbum)
	checkErr(err)






}


func checkErr(err error){
	if err != nil{
		println(err)
	}
}


