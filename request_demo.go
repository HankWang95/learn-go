package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"encoding/json"
	"net/url"
)

type Param struct {
	Word string `json:"word"`
	Translation string `json:"translation"`
}

func main() {


	var param = url.Values{}
	param.Add("word", "100")
	param.Add("translation", "200")
	http.PostForm("http://127.0.0.1:8080/api/add", param)

	return

	methord := http.MethodPost
	url := "http://127.0.0.1:8080/api/add"
	var buf io.Reader
	var j Param

	j.Word = "test"
	j.Translation = "测试"
	data, err := json.Marshal(j)
	if err != nil {
		return
	}
	buf = strings.NewReader(string(data))
	fmt.Println(buf)
	request, err := http.NewRequest(methord, url, buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(request)
	request.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(request)
	fmt.Println(resp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("requset sended!")

}
