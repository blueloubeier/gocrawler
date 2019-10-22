package parser

import (
	"crawler/fetcher"
	"fmt"
	"testing"
)

func TestParseCityList(t *testing.T){
	contents,err :=fetcher.Fetch("http://www.zhenai.com/zhenghun")
	if err !=nil{
		panic(err)
	}
	result :=ParseCityList(contents)
	fmt.Printf("%d\n",len(result.Request))
}
