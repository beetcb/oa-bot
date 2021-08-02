package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func LoadRemoteEnv(pass string) {
	res, err := http.Get(fmt.Sprintf("https://renv.deno.dev/oabot?pass=%s", pass))
	if err != nil {
		log.Fatal("无法获取环境变量信息")
	}
	defer res.Body.Close()
	b, _ := ioutil.ReadAll(res.Body)
	godotenv.Unmarshal(string(b))
}
