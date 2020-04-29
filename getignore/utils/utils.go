package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
func DownloadIfExists(url, filepath string) error {

	// Get the data via HTTP request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Http Error: Cannot Download .gitignore File")
		log.Fatal(err)
		os.Exit(1)
	}

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Open file
	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	// Write file
	_, err = f.WriteString(string(body))
	return err

}

// DownloadIfNotExists downloads contents to a new .gitignore file
func DownloadIfNotExists(url, filepath string) error {

	// Get the data via HTTP request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Http Error: Cannot Download .gitignore File")
		os.Exit(1)
	}

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

	// Create the file if it doesn't exist
	if FileExists(filepath) {
		err := DownloadIfExists(url, filepath)
		if err != nil {
			fmt.Println("Append Error: Cannot Append Content to .gitignore File")
			os.Exit(1)
		}

	} else {
		// Append to previous .gitignore when it already exists
		err := DownloadIfNotExists(url, filepath)
		if err != nil {
			fmt.Println("Add Error: Cannot Add Content to .gitignore File")
			os.Exit(1)
		}
	}
}
