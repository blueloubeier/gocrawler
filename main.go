package main

import (
	"crawler/engine"
	"crawler/zhenai/parser"
)

func main() {

	request := engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	}

	engine.Run(request)
}