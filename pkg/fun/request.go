package fun

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

// NewGetRequest 创建一个get请求
func NewGetRequest(u string, q map[string]string, h map[string]string) (*http.Request, error) {
	if q != nil {
		qs := url.Values{}

		for k, v := range q {
			qs.Set(k, v)
		}

		u = u + "?" + qs.Encode()
	}

	req, err := http.NewRequest("GET", u, nil)

	if err != nil {
		return nil, err
	}

	if h != nil {
		for k, v := range h {
			req.Header.Set(k, v)
		}
	}

	return req, nil
}

// NewPostRequest 创建一个post请求
func NewPostRequest(u string, p map[string]string, h map[string]string) (*http.Request, error) {

	var reader *strings.Reader

	if p != nil {
		qs := url.Values{}

		for k, v := range p {
			qs.Set(k, v)
		}

		reader = strings.NewReader(qs.Encode())
	}

	req, err := http.NewRequest("POST", u, reader)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if h != nil {
		for k, v := range h {
			req.Header.Set(k, v)
		}
	}

	return req, nil
}

// NewJsonRequest 创建一个application/json请求
func NewJsonRequest(u string, j []byte, h map[string]string) (*http.Request, error) {
	req, err := http.NewRequest("POST", u, bytes.NewReader(j))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	if h != nil {
		for k, v := range h {
			req.Header.Set(k, v)
		}
	}

	return req, nil
}

// NewFilesRequest 多文件上传
// u url 请求的api路径
// fs 上传的文件  eg: [][]string{{"fieldname","filepath","filename"}}
// p 上传的普通post参数
// h header参数
func NewFilesRequest(u string, fs [][]string, p map[string]string, h map[string]string) (*http.Request, error) {

	if len(fs) == 0 {
		return nil, errors.New("文件参数为空")
	}

	//检查文件是否存在
	for _, item := range fs {

		if len(item) < 2 {
			return nil, errors.New("文件参数格式错误")
		}

		_, err := os.Lstat(item[1])

		if os.IsNotExist(err) {
			return nil, err
		}
	}

	readerBuf := new(bytes.Buffer)

	writer := multipart.NewWriter(readerBuf)

	for _, item := range fs {

		var filename string

		if len(item) >= 3 {
			filename = item[2]
		} else {
			filename = path.Base(item[1])
		}

		fileWriter, err := writer.CreateFormFile(item[0], filename)

		if err != nil {
			return nil, err
		}

		file, err := os.Open(item[1])

		if err != nil {
			return nil, err
		}

		io.Copy(fileWriter, file)

		//关闭文件
		if err = file.Close(); err != nil {
			return nil, err
		}
	}

	writer.Close()

	if p != nil {

		for k, v := range p {
			if err := writer.WriteField(k, v); err != nil {
				return nil, err
			}
		}

	}

	req, err := http.NewRequest("POST", u, readerBuf)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	if h != nil {
		for k, v := range h {
			req.Header.Set(k, v)
		}
	}

	return req, nil
}
