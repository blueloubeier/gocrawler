package parser

import (
	"crawler/fetcher"
	"fmt"
	"testing"
)

func TestParseCityList(t *testing.T){
	contents,err :=fetcher.Fetch("http://www.zhenai.com/zhenghun/aba")
	if err !=nil{
		panic(err)
	}
	result :=ParseCity(contents)
	fmt.Printf("%d\n",len(result.Request))
}
