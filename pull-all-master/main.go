package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.Name() == ".git" {
				path := strings.Replace(dir+"/"+path, info.Name(), "", 1)

				_, err := exec.Command("git", "-C", path, "checkout", "master").
					Output()
				if err != nil {
					log.Fatal(err)
				}

				_, err = exec.Command("git", "-C", path, "pull").
					Output()
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("Successfully updated repo: ", path)
			}
			return nil
		})
}
