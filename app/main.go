package main

import (
	t "app/tracker"

	"log"
	"path/filepath"
)

func main() {
	absPath, err := filepath.Abs(".env")
	if err != nil {
		log.Fatalln("[!] ", err)
		return
	}

	t.Run(absPath)
}
