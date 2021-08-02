package extract

import (
	"fmt"
	"os"

	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
)

func ExtractPdfText(inputPath string) string {

	err := license.SetMeteredKey(os.Getenv(`UNIDOC_LICENSE_API_KEY`))
	if err != nil {
		panic(err)
	}

	f, _ := os.Open(inputPath)
	fmt.Println(inputPath)

	defer f.Close()

	pdfReader, err := model.NewPdfReader(f)
	if err != nil {
		panic(err)
	}

	isEncrypted, err := pdfReader.IsEncrypted()
	if err != nil {
		panic(err)
	}

	if isEncrypted {
		_, err := pdfReader.Decrypt([]byte(""))
		if err != nil {
			panic(err)
		}
	}

	pdfWriter, _ := pdfReader.ToWriter(nil)

	pdfWriter.WriteToFile(inputPath)

	p, err := pdfReader.GetPage(1)
	if err != nil {
		panic(err)
	}
	ex, _ := extractor.New(p)
	text, _ := ex.ExtractText()

	return text
}
