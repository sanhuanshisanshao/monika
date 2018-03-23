package util

import (
	"fmt"
	"testing"
)

func TestSplit(t *testing.T) {
	str := "03月12日 10:33 来自iPhone客户端"
	sep := "来自"
	s := Split(sep, str)
	if len(s) != 2 {
		t.Fatalf("split length is not equal two")
	}
	fmt.Println(s)

	res, err := ConvertTime(s[0])
	if err != nil {
		t.Fatalf("ConvertTime error %v", err)
	}
	fmt.Println(res)

	str = "2014-06-19 19:19:50 来自Windows.Phone客户端"
	s = Split(sep, str)
	if len(s) != 2 {
		t.Fatalf("split length is not equal two")
	}
	fmt.Println(s)

	res, err = ConvertTime(s[0])
	if err != nil {
		t.Fatalf("ConvertTime error %v", err)
	}
	fmt.Println(res)
}
