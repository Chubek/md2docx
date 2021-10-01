package util

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"github.com/joho/godotenv"

)


func DownloadFile(filepath string, url string) (string, error) {
	err_env := godotenv.Load()
	if err_env != nil {
		log.Fatal("Error loading .env file")
	}

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	imgSavePath := path.Join(os.Getenv(`FILES_PATH`), filepath)

	out, err := os.Create(imgSavePath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return imgSavePath, err
}

func Reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}