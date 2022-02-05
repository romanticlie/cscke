package fun

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var (
	Y        = "2006"
	M        = "01"
	D        = "02"
	H        = "15"
	I        = "04"
	S        = "05"
	DateTime = fmt.Sprintf("%s-%s-%s %s:%s:%s", Y, M, D, H, I, S)
	seedOnce sync.Once
)

// GetYamlCfg 获取cfg的yaml文件配置
func GetYamlCfg(filename string) (*viper.Viper, error) {

	//获取项目的执行路径
	v := viper.New()
	v.AddConfigPath("/www/pkg/conf")
	v.SetConfigType("yaml")
	v.SetConfigName(filename)

	//尝试读取文件
	if err := v.ReadInConfig(); err != nil {
		return v, err
	}

	return v, nil
}

// Random 随机一个范围的数
func Random(min int, max int) int {

	seedOnce.Do(func() {
		rand.Seed(time.Now().UnixNano())
	})

	max = max + 1

	return rand.Intn(max-min) + min
}

// MapStringKeys 获取map的key值
func MapStringKeys(m map[string]string) []string {

	keys := make([]string, len(m))

	i := 0

	for k, _ := range m {
		keys[i] = k
		i++
	}

	return keys
}

// HttpGet 普通的GET请求
func HttpGet(u string, q map[string]string, h map[string]string) (b []byte, err error) {

	client := &http.Client{}

	req, err := NewGetRequest(u, q, h)

	if err != nil {
		return
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

	client := &http.Client{}

	req, err := NewPostRequest(u, p, h)

	if err != nil {
		return
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

	req, err := NewJsonRequest(u, j, h)

	if err != nil {
		return
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

	client := &http.Client{}

	req, err := NewFilesRequest(u, fs, p, h)

	if err != nil {
		return
	}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return
	}

	return ioutil.ReadAll(resp.Body)
}
