package test

import "testing"

func TestCopy(t *testing.T){

	s := [][]int{
		{1,2,3},
		{4,5,6},
	}

	b := make([][]int,2,2)

	copy(b,s)

	s[0][0] = 2

	t.Log(b)

}
