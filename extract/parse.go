package extract

import (
	"path/filepath"
	"regexp"
	"strings"
)

type ParseInfo struct {
	FileName string
	Number   string
	Date     string
}

func ParsePdfText(inputPath, text string) ParseInfo {
	inputPath = filepath.ToSlash(inputPath)

	re := regexp.MustCompile(`([0-9_]){6,10}`)
	fileName := re.ReplaceAllString(filepath.Base(inputPath), "")
	fileName = strings.Replace(fileName, ".pdf", "", 1)

	re = regexp.MustCompile("[\u4e00-\u9fa5]+〔\\d{4}〕\\d+ 号")
	number := re.FindString(text)

	re = regexp.MustCompile(`(\d{4})\s年\s(\d{1,2})\s月\s(\d{1,2})\s日`)
	dateParts := re.FindStringSubmatch(text)
	var date string
	if len(dateParts) > 0 {
		date = strings.Join(re.FindStringSubmatch(text)[1:], "-")
	}

	return ParseInfo{fileName, number, date}

}
