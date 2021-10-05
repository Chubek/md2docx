package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "os"

	_ "github.com/unidoc/unioffice/common/license"
)


func init() {
	// Make sure to load your metered License API key prior to using the library.
	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io
	err := license.SetMeteredKey(os.Getenv(`UNIDOC_LICENSE_API_KEY`))
	if err != nil {
		panic(err)
	}
}
func main() {
	
}
