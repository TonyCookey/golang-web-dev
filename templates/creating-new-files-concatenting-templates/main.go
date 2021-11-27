package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Tony Cookey"

	webPage := `
			<!DOCTYPE html>
			<html lang="en">
			<head>
			<meta charset ="UTF-8">
			<title>Hello World</title>
			</head>
			<body>
			<h1>` + name + `</h1>
			</body>
			</html>`
	newFile, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, strings.NewReader(webPage))
	if err != nil {
		log.Fatalln(err)
	}
}
