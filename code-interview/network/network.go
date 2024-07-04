package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// DownloadToFile downloads a given URL to the target filepath
func DownloadToFile(src string, destinationFile string, cosignKeyPath string) (err error) {

	WhereAmI(destinationFile)

	// Create the file
	file, err := os.Create(destinationFile)
	if err != nil {
		return fmt.Errorf("error creating file: %s, %s", destinationFile, err.Error())
	}

	HttpGetFile(src, file)

	return nil
}

func WhereAmI(dest string) {
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
	_ = os.Chdir(dest)
}

func HttpGetFile(url string, destinationFile *os.File) {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	if _, err = io.Copy(destinationFile, resp.Body); err != nil {
		panic(err)
	}

}
