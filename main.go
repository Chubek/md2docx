package main

import (
	"flag"
	"fmt"
	"md2docx/process"
)

func main() {
	mdFilePtr := flag.String("mdFile", "", "Path to Matkdown file")
	savePathPtr := flag.String("savePath", "", "Path to save file")

	fmt.Printf("Loaded Markdown file %s and save path %s\n", *mdFilePtr, *savePathPtr)

	fmt.Println("Converting Markdown to Word...")

	process.ProcessAndWrite(*mdFilePtr, *savePathPtr)

	fmt.Println("Done! Exiting. If you like this piece of software, contact it's author at chubakbidpaa.com!")

}
