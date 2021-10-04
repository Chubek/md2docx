package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "os"

	_ "github.com/unidoc/unioffice/common/license"
)

/*
func init() {
	// Make sure to load your metered License API key prior to using the library.
	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io
	err := license.SetMeteredKey(os.Getenv(`UNIDOC_LICENSE_API_KEY`))
	if err != nil {
		panic(err)
	}
}*/
func main() {
	resp, err := http.Get("https://media.discordapp.net/attachments/797653922894184488/894606892515618866/Screenshot_2021-10-04-16-27-43-59_5bc87d1a35b644312bd0b78de3cfc1ea.jpg?width=213&height=473")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(http.DetectContentType(bodyBytes))
}
