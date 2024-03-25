package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	res, err := sha1Sum("sha1.go")
	if err != nil {
		log.Fatalf("Error%s", err)
	}
	fmt.Println(res)
}

// $ cat http.log.gz | sha1sum
func sha1Sum(fileName string) (string, error) {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Unable to open the file %s", err)
	}

	defer file.Close()
	var r io.Reader = file

	if strings.HasSuffix(fileName, ".gz") {
		gz, err := gzip.NewReader(file)
		if err != nil {
			log.Fatal("Error in gzip")
		}
		defer gz.Close()
		r = gz
	}

	w := sha1.New()

	if _, err := io.Copy(w, r); err != nil {
		log.Fatal("Error in io")
	}

	sig := w.Sum(nil)
	return fmt.Sprintf("%x", sig), nil
}
