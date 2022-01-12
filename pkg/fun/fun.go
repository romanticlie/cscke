package fun

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"
)

var (
	Y = "2006"
	M = "01"
	D = "02"
	H = "15"
	I = "04"
	S = "05"
	DateTime = fmt.Sprintf("%s-%s-%s %s:%s:%s",Y,M,D,H,I,S)
)

// GetYamlCfg 获取cfg的yaml文件配置
func GetYamlCfg(filename string) (*viper.Viper,error){

	//获取项目的执行路径
	v := viper.New()
	v.AddConfigPath("/www/pkg/conf")
	v.SetConfigType("yaml")
	v.SetConfigName(filename)

	//尝试读取文件
	if err := v.ReadInConfig(); err != nil {
		return v,err
	}

	return v,nil
}


// Random 随机一个范围的数
func Random(min int,max int) int{

	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max - min) + min
}

// MapStringKeys 获取map的key值
func MapStringKeys(m map[string]string) []string{

	keys := make([]string,len(m))

	i := 0

	for k,_ := range m {
		keys[i] = k
		i++
	}

	return keys
}



// HttpGet 普通的GET请求
func HttpGet(u string, q map[string]string, h map[string]string) (b []byte, err error) {

	if q != nil {
		qs := url.Values{}

		for k, v := range q {
			qs.Set(k, v)
		}

		u = u + "?" + qs.Encode()
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", u, nil)

	if err != nil {
		return
	}

	if h != nil {
		for k, v := range h {
			req.Header.Set(k, v)
		}
	}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return
	}

	return ioutil.ReadAll(resp.Body)
}


// HttpPost 普通的POST请求
func HttpPost(u string, p map[string]string, h map[string]string) (b []byte, err error) {

	var reader *strings.Reader

	if p != nil {
		qs := url.Values{}

		for k, v := range p {
			qs.Set(k, v)
		}

		reader = strings.NewReader(qs.Encode())
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", u, reader)

	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if h != nil {
		for k, v := range h {
			req.Header.Set(k, v)
		}
	}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return
	}

	return ioutil.ReadAll(resp.Body)
}


// HttpJson JSON的POST请求
func HttpJson(u string, j []byte, h map[string]string) (b []byte, err error) {

	client := &http.Client{}

	req, err := http.NewRequest("POST", u, bytes.NewReader(j))

	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")

	if h != nil {
		for k, v := range h {
			req.Header.Set(k, v)
		}
	}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return
	}

	return ioutil.ReadAll(resp.Body)
}

// HttpFiles 多文件上传+
// u url 请求的api路径
// fs 上传的文件  eg: [][]string{{"fieldname","filepath","filename"}}
// p 上传的普通post参数
// h header参数
func HttpFiles(u string, fs [][]string, p map[string]string, h map[string]string) (b []byte, err error) {

	if len(fs) == 0 {
		return
	}

	//检查文件是否存在
	for _, item := range fs {

		if len(item) < 2 {
			return
		}

		_, err := os.Lstat(item[1])

		if os.IsNotExist(err) {
			return b, err
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
			return b, err
		}

		file, err := os.Open(item[1])

		if err != nil {
			return b, err
		}

		io.Copy(fileWriter, file)

		//关闭文件
		if err := file.Close(); err != nil {
			return b, err
		}
	}

	writer.Close()

	if p != nil {

		for k, v := range p {
			if err := writer.WriteField(k, v); err != nil {
				return b, err
			}
		}

	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", u, readerBuf)

	if err != nil {
		return
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	if h != nil {
		for k, v := range h {
			req.Header.Set(k, v)
		}
	}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return
	}

	return ioutil.ReadAll(resp.Body)
}


