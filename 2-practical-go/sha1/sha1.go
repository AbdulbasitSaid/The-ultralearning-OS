package main

import (
	"compress/gzip"
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println(SHA1Sig("http.log.gz"))
}

// SHA1Sig return SHA1 signature of uncompressed file
// > cat http.log.gz | gunzip | sha1sum
func SHA1Sig(fileName string) (string, error) {
	if strings.HasSuffix(fileName, ".gz") {
		// cat http.log.gz
		file, err := os.Open(fileName)

		if err != nil {
			return "", fmt.Errorf("%q - error reading file %w", fileName, err)
		}
		defer file.Close()

		// | gunzip
		reader, err := gzip.NewReader(file)

		if err != nil {
			return "", fmt.Errorf("%q - gzip: %w", fileName, err)
		}
		defer reader.Close()

		// | sha1sum
		writer := sha1.New()
		if _, err := io.Copy(writer, reader); err != nil {
			return "", fmt.Errorf("%q - copy: %w", fileName, err)

		}
		sig := writer.Sum(nil)

		return fmt.Sprintf("%x", sig), nil
	} else {
		return " ", errors.New("function expects only .gz file type")
	}

}
