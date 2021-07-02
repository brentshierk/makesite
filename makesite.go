package main

import (
	"flag"
	"fmt"
	//"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"context"
	//"path/filepath"
	
	
	"strings"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)



func main() {
	var fileName,directory string

	flag.StringVar(&fileName,"file","","text file you want to turn into an html page")
	flag.StringVar(&directory,"directory",".","directory to look for text files to convert (all)")
	flag.Parse()

	

	if directory != ""{
		dirToHtml(directory)
	}else if fileName != ""{
		makePlaylist(fileName)
		fmt.Println("this function ran")
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

func makePlaylist(fileName string){
	fmt.Println("function running")
	config := &clientcredentials.Config{
		ClientID : "32d484634e5843ee80024a0d65d1459d",
		ClientSecret : "4629c73660b8499fa4809255e4bba77e",
		TokenURL: spotify.TokenURL,
	}


	token, err := config.Token(context.Background())
	checkErr(err)

client := spotify.Authenticator{}.NewClient(token)

fileContent, err := ioutil.ReadFile(fileName)
	checkErr(err)
text := string(fileContent)
fmt.Println(text)
q := strings.SplitN(text,",",5)
fmt.Println(q)

for _,i  := range q{
	results, err := client.Search(i, spotify.SearchTypePlaylist|spotify.SearchTypeAlbum)
	
	checkErr(err)

	// handle album results
	if results.Albums != nil {
		fmt.Println("Albums:")
		for _, item := range results.Albums.Albums {
			fmt.Println(" ", item.Name)
		}
	}
	// handle playlist results
	if results.Playlists != nil {
		fmt.Println("Playlists:")
		for _, item := range results.Playlists.Playlists {
			fmt.Println(" ", item.Name)
		}
	}
}
//split the content of the file at ","


// for each substring compleate a query to the spotify api

}
//
//return json of playlist
//get  playlist icon,title,link to playlist










func checkErr(err error){
	if err != nil{
		println(err)
	}
}


