package request

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	JSON       = "application/json"
	URLENCODED = "application/x-www-form-urlencoded"
)

func PostFormData() error {
	body := make(url.Values)

	body.Add("name", "lxc")

	req, err := http.NewRequest("POST", "http://127.0.0.1:3500/post", strings.NewReader(body.Encode()))
	if err != nil {
		return err
	}

	// 设置header
	// 指定contentType
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	do, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer do.Body.Close()

	io.Copy(os.Stdout, do.Body)

	return nil
}

func PostJsonData() error {
	body := struct {
		Name string `json:"name"`
	}{
		Name: "lxc",
	}
	jsonBody, err := json.Marshal(&body)
	if err != nil {
		return err
	}
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest("POST", "http://127.0.0.1:3500/post", bodyReader)
	if err != nil {
		return err
	}

	// 设置头信息
	// 设置content-type
	req.Header.Set("Content-Type", "application/json")

	do, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer do.Body.Close()

	io.Copy(os.Stdout, do.Body)

	return nil
}

func UploadFile() error {
	var body bytes.Buffer

	writer := multipart.NewWriter(&body)
	if err := writer.WriteField("name", "liuxuech"); err != nil {
		return err
	}

	// 文件1
	fileWriter1, err := writer.CreateFormFile("file1", "README.md")
	if err != nil {
		return err
	}
	fd, err := os.Open("README.md")
	if err != nil {
		return err
	}
	if _, err := io.Copy(fileWriter1, fd); err != nil {
		return err
	}

	// 文件2
	fileWriter2, err := writer.CreateFormFile("file2", "README.md")
	fd, err = os.Open("README.md")
	if err != nil {
		return err
	}
	if _, err := io.Copy(fileWriter2, fd); err != nil {
		return err
	}

	if err := writer.Close(); err != nil {
		return err
	}

	contentType := writer.FormDataContentType()

	resp, err := http.Post("http://127.0.0.1:3500/upload", contentType, &body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)

	return nil
}

// 这里传url.Values没有传string，考虑的时候可以添加一些默认的请求参数
func Get(url string, query url.Values) (err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// 设置请求参数
	req.URL.RawQuery = query.Encode()

	//resp, err := http.DefaultClient.Do(req)
	//if err != nil {
	//	return err
	//}
	//
	//http.Transport{}
	return nil
}
