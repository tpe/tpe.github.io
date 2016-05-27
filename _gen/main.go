package main

import (
	"fmt"
    "os"
	"path"
    "strings"
)

var posts Posts

// This process cleans the directory and then creates the pages.
func main() {
	fmt.Println("*** START ***")

    pathGen, _ := os.Getwd()
    pathRoo := strings.Replace(pathGen, "/_gen", "", -1)
	pathSrc := path.Join(pathRoo, "/_src")
	pathPst := path.Join(pathSrc, "/posts")
	pathTmp := path.Join(pathSrc, "/templates")

	fileTmp := GetFile(path.Join(pathTmp, "post.html"))

	ProcessClean(pathRoo)
	ProcessPosts(fileTmp, pathPst, pathRoo)

	fmt.Println("*** END ***")
}
