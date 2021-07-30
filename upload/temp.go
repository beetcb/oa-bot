package upload

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func UploadToTemp(inputPath string) string {
	fileBytes, err := ioutil.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	reqs, _ := http.NewRequest("PUT", "https://transfer.sh/FILE.pdf", strings.NewReader(string(fileBytes)))

	reqs.Header.Set("Max-Downloads", "1")
	res, err := http.DefaultClient.Do(reqs)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	b, _ := ioutil.ReadAll(res.Body)
	url := string(b)

	return url
}
