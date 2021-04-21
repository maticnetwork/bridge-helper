package main

import (
	"log"
	"path/filepath"
	app "state-id-manager/app"
)

func main() {
	path, err := filepath.Abs(".env")
	if err != nil {
		log.Fatalln("[!] ", err)
		return
	}

	app.Run(path)
}
