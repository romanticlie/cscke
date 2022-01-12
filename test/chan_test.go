package test

import (
	"cscke/pkg/fun"
	"testing"
)

func TestChan(t *testing.T){

	var (
		Baidu = "baidu"
		GinGuess = "ginGuess"
	)

	requestCh := make(chan map[string]string,2)


	go func() {
		b,err := fun.HttpGet("https://www.baidu.com",nil,nil)

		errMsg := ""

		if err != nil {
			errMsg = err.Error()
		}

		requestCh <- map[string]string{
			"type": Baidu,
			"ret": string(b),
			"errMsg": errMsg,
		}
	}()


	go func() {
		b,err := fun.HttpGet("http://ginguess.com",nil,nil)

		errMsg := ""

		if err != nil {
			errMsg = err.Error()
		}

		requestCh <- map[string]string{
			"type": GinGuess,
			"ret": string(b),
			"err": errMsg,
		}
	}()

	var receives = make([]map[string]string,0,2)

	for i := 0;i < 2;i++ {

		select {
		case j := <- requestCh:
			receives = append(receives,j)
		}
	}

	t.Log(receives)
}

func TestChanBenchmark(t *testing.T){

	count := int(1e6)
	ch := make(chan int,count)

	for i := 0;i < count;i++ {
		go func(ch chan int) {
			ch <- 1
		}(ch)
	}

	s := make([]int,0,count)
	for i := 0; i < count;i++ {
		select {
		case j := <- ch:
			s = append(s,j)
		}
	}

	t.Log(len(s))
}

