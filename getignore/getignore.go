package main

import (
	"fmt"
	"getIgnore/getignore/langs"
	"getIgnore/getignore/utils"
	"github.com/urfave/cli/v2" // imports as package "cli"
	"log"
	"os"
	"os/exec"
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
		fmt.Printf("Gitignore for %s not found\n", lang)
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
				Usage:       "Provide the desired list of languages",
				Destination: &languages,
			},
			&cli.BoolFlag{Name: "list", Aliases: []string{"ls"}},
		},

		Action: func(c *cli.Context) error {
			if c.NArg() == 0 {
				base := "getignore"
				arg0 := "-h"

				cmd := exec.Command(base, arg0)
				stdout, err := cmd.Output()
				if err != nil {
					fmt.Println(err.Error())
					return err
				}
				fmt.Print(string(stdout))
				os.Exit(0)
			}

			if os.Args[1] == "--lg" || os.Args[1] == "--languages" {
				for _, lang := range os.Args[2:] {
					langURL := SelectLang(langMap, lang)
					if langURL != "" {
						utils.DownloadFile(langURL, "./.gitignore")
						fmt.Printf("Downloading %s gitignore\n", strings.Title(lang))

					}
					cli.Exit("", 0)
				}
			}

			if os.Args[1] == "--ls" || os.Args[1] == "--list" {
				fmt.Println("Language List")
				fmt.Println("===============")

				for _, lang := range langList {
					fmt.Println(lang)
					cli.Exit("", 0)
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
