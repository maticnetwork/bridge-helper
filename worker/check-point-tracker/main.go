package main

import (
	"check-point-tracker/app"
	"log"
	"path/filepath"
)

func main() {
	path, err := filepath.Abs(".env")
	if err != nil {
		log.Fatalln("[!] ", err)
		return
	}

	app.Run(path)
}
