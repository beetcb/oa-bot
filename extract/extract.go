package extract

import (
	"os"

	"github.com/unidoc/unipdf-cli/pkg/pdf"
	"github.com/unidoc/unipdf/v3/common/license"
)

func FillLicense() {
	err := license.SetMeteredKey(os.Getenv(`UNIDOC_LICENSE_API_KEY`))
	if err != nil {
		panic(err)
	}
}

func ExtractPdfText(inputPath string) string {
	text, err := pdf.ExtractText(inputPath, "", nil)
	if err != nil {
		panic(err)
	}
	return text
}
