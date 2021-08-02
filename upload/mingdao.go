package upload

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/beetcb/oa-bot/extract"
)

type Controls struct {
	ControlId string `json:"controlId"`
	Value     string `json:"value"`
}

type UploadBody struct {
	AppKey      string     `json:"appKey"`
	Sign        string     `json:"sign"`
	WorksheetId string     `json:"worksheetId"`
	Controls    []Controls `json:"controls"`
}

func UploadToMingDao(tempUrl string, parse extract.ParseInfo) string {
	uploadBody := UploadBody{
		AppKey:      os.Getenv(`APP_KEY`),
		Sign:        os.Getenv(`SIGN`),
		WorksheetId: os.Getenv(`WORKSHEET_ID`),
		Controls: []Controls{
			{ControlId: "file", Value: tempUrl},
			{
				ControlId: "number",
				Value:     parse.Number,
			}, {
				ControlId: "date",
				Value:     parse.Date,
			}, {
				ControlId: "title",
				Value:     parse.FileName,
			}},
	}
	uploadBodyJson, _ := json.Marshal(uploadBody)
	res, err := http.Post("https://api.mingdao.com/v2/open/worksheet/addRow", "application/json", strings.NewReader(string(uploadBodyJson)))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	b, _ := ioutil.ReadAll(res.Body)
	return fmt.Sprintf("明道云上传结果：%s", string(b))
}
