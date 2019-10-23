package parser

import (
	"crawler/fetcher"
	"fmt"
	"testing"
)

func TestParseCityList(t *testing.T){
	contents,err :=fetcher.Fetch("https://album.zhenai.com/u/1403643364")
	if err !=nil{
		panic(err)
	}
	result :=ParseProfile(contents,"james")
	fmt.Printf("%+v\n" ,result)
}
