package main

import (
	"flag"
	"fmt"

	//"fmt"
	"context"
	"html/template"
	//"io/fs"
	"io/ioutil"
	"os"

	//"path/filepath"
	"bufio"

	"strings"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)



func main() {
	var fileName,directory string

	flag.StringVar(&fileName,"file","","text file you want to turn into an html page")
	flag.StringVar(&directory,"directory",".","directory to look for text files to convert (all)")
	flag.Parse()

	fmt.Println("hello world")
	//makePlaylist("test.txt")
	if directory != ""{
		dirToHtml(directory)
	}else if fileName != ""{
		makePlaylist(fileName)
		fmt.Println("this function ran")
	}

	
}

func fileToHtml(fileName string){
	makePlaylist(fileName)
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
	//the tmp file will be  used later to create the template,
	//but first the results of the spotify api query need to be writen to a text file.
	tmp,err:= os.Create(fileName +"(albums).txt")
	checkErr(err)
	println(tmp)
	tmp2,err:= os.Create(fileName +"(playlist).txt")
	println(tmp2)
	checkErr(err)

	//println("file created for "+fileName +"(result).txt")
	
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
	
	q := strings.SplitN(text,",",5)
	
	
	for _,i := range q {
	results, err := client.Search(i, spotify.SearchTypePlaylist|spotify.SearchTypeAlbum)
	checkErr(err)
	var s []string
	// handle album results
	if results.Albums != nil {
		fmt.Println("Albums:")
		for _, item := range results.Albums.Albums {
			fmt.Println("   ", item.Name)
			s = append(s, item.Name)
			
		}
		//var hold []byte
		//hold = append(hold,s)
		writer := bufio.NewWriter(tmp)
		var j int 
		for j=0;j<10;j++ {
			_, err := writer.WriteString(s[j] + "\n")
			checkErr(err)
			writer.Flush()
	}
	// handle playlist results
	if results.Playlists != nil {
		fmt.Println("Playlists:")
		for _, item := range results.Playlists.Playlists {
			fmt.Println("   ", item.Name)
		}
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

