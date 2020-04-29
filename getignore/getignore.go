package main

import (
	"fmt"
	"getIgnore/getignore/langs"
	"getIgnore/getignore/utils"
	"github.com/urfave/cli/v2" // imports as package "cli"
	"log"
	"os"
	"strings"
)

// URLMap will map language: url
func URLMap(langList []string) map[string]string {
	var langMap map[string]string
	langMap = make(map[string]string)

	p1 := "https://raw.githubusercontent.com/"

	for _, lang := range langList {
		p2 := fmt.Sprintf("github/gitignore/master/%s.gitignore", lang)
		url := p1 + p2
		langMap[lang] = url
	}
	return langMap

}

// SelectLang will select the url from cli argument
func SelectLang(langMap map[string]string, lang string) string {
	lang = strings.Title(lang)
	langURL := langMap[lang]
	if langURL == "" {
		fmt.Printf("Gitignore for %s not found.\n", lang)
		os.Exit(1)
	}
	return langURL
}

// MakeCli gives the final CLI
func MakeCli() {
	var languages string
	langList := langs.GetLangs()
	langMap := URLMap(langList)

	app := &cli.App{
		Name:  "getignore",
		Usage: "Download Gitignore Files",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "languages",
				Aliases:     []string{"lg"},
				Usage:       "Prints the names of the supported languages",
				Destination: &languages,
			},
		},
		Action: func(c *cli.Context) error {
			if os.Args[1] == "--lg" || os.Args[1] == "--languages" {
				for _, lang := range os.Args[2:] {
					langURL := SelectLang(langMap, lang)
					fmt.Printf("Downloading %s gitignore\n", lang)
					utils.DownloadFile(langURL, "./.gitignore")

				}
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	MakeCli()

}
