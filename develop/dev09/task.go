package main

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/pborman/getopt"
)

func main() {
	file := getopt.StringLong("url", 'u', "", "utl")
	getopt.Parse()
	fileSplit := strings.Split(*file, "/")
	filepath := strings.Split(fileSplit[len(fileSplit)-1], ".")
	filepathSplit := filepath[0]
	err := downloadFile(filepathSplit, *file, filepath[1])

	if err != nil {
		log.Fatalln(err)
	}

}
func downloadFile(filepath string, url string, addition string) (err error) {
	err = os.Mkdir(filepath, 0777)
	if err != nil {
		return err
	}
	// Create the file
	out, err := os.Create(filepath + "/" + filepath + "." + addition)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
