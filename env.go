package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func LoadRemoteEnv() {
	res, err := http.Get(os.Getenv("REMOTE_ENV"))
	if err != nil {
		log.Fatal("无法获取环境变量信息")
	}
	defer res.Body.Close()
	b, _ := ioutil.ReadAll(res.Body)
	godotenv.Unmarshal(string(b))
}
