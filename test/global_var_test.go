package test

import (
	"cscke/pkg/policy"
	"sync"
	"testing"
)


func TestBuildUrl(t *testing.T){

	p := &policy.WechatWebStrategy{}

	t.Log(p.BuildAuthUrl("STATE"))
}

func TestConfig(t *testing.T){
	p := &policy.WechatWebStrategy{}

	var w sync.WaitGroup

	w.Add(10)

	for i := 0;i < 10;i++{
		go func() {
			p.GetConfig()
			w.Done()
		}()
	}


	w.Wait()

	t.Log(p.GetConfig())
}


func TestAssignment(t *testing.T){
	var a int

	var w sync.WaitGroup
	count := int(1e7)

	w.Add(count)

	for i := 0; i < count; i++ {
		go func(i int) {

			a = i

			w.Done()
		}(i)
	}


	w.Wait()

	t.Log(a)
}


func TestMap(t *testing.T){

	var m sync.Map

	m.Store("name","san")
	m.Store("age",18)
	m.Store("weight",120)

	v,_ := m.Load("name")

	t.Log(v)

	m.Range(func(key, value interface{}) bool {
		t.Log(key,value)

		return true
	})

	m.Delete("name")


	m.Range(func(key, value interface{}) bool {
		t.Log(key,value)

		return true
	})


}

func TestSlice(t *testing.T){

	s := []int{0,1,2,3,4,5,6,7}

	s1 := s[1:3:6]

	t.Log(s1)
	t.Log(len(s1),cap(s1))
}