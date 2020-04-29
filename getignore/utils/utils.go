package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"log"
)

// FileExists checks if a file exists and is not a directory
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// DownloadIfExists appends content when .gitignore exists
func DownloadIfExists(resp *http.Response, filepath string) error {
	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Open File
	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	// Write File
	_, err = f.WriteString(string(body))
	return err

}

// DownloadIfNotExists downloads contents to a new .gitignore file
func DownloadIfNotExists(resp *http.Response, filepath string) error {

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// DownloadFile from url
func DownloadFile(url string, filepath string) {

	// Get the data via HTTP request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Http Error: Cannot Download .gitignore File")
		log.Fatal(err)
		os.Exit(1)
	}

	// Create the file if it doesn't exist
	if FileExists(filepath) {
		err := DownloadIfExists(resp, filepath)
		if err != nil {
			fmt.Println("Append Error: Cannot Append Content to .gitignore File")
			os.Exit(1)
		}

	} else {
		// Append to previous .gitignore when it already exists
		err := DownloadIfNotExists(resp, filepath)
		if err != nil {
			fmt.Println("Add Error: Cannot Add Content to .gitignore File")
			os.Exit(1)
		}
	}
}
